package mtproto

import (
	"fmt"
	"math/big"
	"math/rand"
	"net"
	"os"
	"time"
)

const (
	appId   = 2899
	appHash = "36722c72256a24c1225de00eb6a1ca74"
)

type MTProto struct {
	conn      *net.TCPConn
	f         *os.File
	queueSend chan packetToSend

	authKey     []byte
	authKeyHash []byte
	serverSalt  []byte
	encrypted   bool
	sessionId   int64

	lastSeqNo    int32
	msgsIdToAck  map[int64]bool
	msgsIdToResp map[int64]chan TL
	seqNo        int32
	msgId        int64

	dclist map[int32]string
}

type packetToSend struct {
	msg  interface{}
	resp chan TL
}

func NewMTProto(addr, authkeyfile string) (*MTProto, error) {
	var err error
	var tcpAddr *net.TCPAddr

	m := new(MTProto)

	// try to read [authKey, serverSalt]
	// TODO: read [dcAddr, authStatus]
	m.f, err = os.OpenFile(authkeyfile, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}
	b := make([]byte, 256+8+8)
	n, err := m.f.Read(b)
	if n == 256+8+8 {
		m.authKey = b[:256]
		m.authKeyHash = b[256 : 256+8]
		m.serverSalt = b[256+8:]
		m.encrypted = true
	} else {
		m.encrypted = false
	}
	rand.Seed(time.Now().UnixNano())
	m.sessionId = rand.Int63()

	// connect
	tcpAddr, err = net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	m.conn, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, err
	}
	_, err = m.conn.Write([]byte{0xef})
	if err != nil {
		return nil, err
	}

	// get new authKey if need
	if !m.encrypted {
		err = m.makeAuthKey()
		if err != nil {
			return nil, err
		}
	}

	// start goroutines
	m.queueSend = make(chan packetToSend, 64)
	m.msgsIdToAck = make(map[int64]bool)
	m.msgsIdToResp = make(map[int64]chan TL)
	go m.SendRoutine()
	go m.ReadRoutine()

	var resp chan TL
	var x TL

	// (help_getConfig)
	resp = make(chan TL, 1)
	m.queueSend <- packetToSend{
		&TL_invokeWithLayer{
			layer,
			&TL_initConnection{
				appId,
				"MacBook Pro (Retina, Mid 2012)",
				"OS X 10.10.3",
				"0.0.1",
				"en",
				&TL_help_getConfig{},
			},
		},
		resp,
	}
	x = <-resp
	switch x.(type) {
	case *TL_config:
		m.dclist = make(map[int32]string, 5)
		for _, v := range x.(*TL_config).dc_options {
			v := v.(*TL_dcOption)
			m.dclist[v.id] = fmt.Sprintf("%s:%d", v.ip_address, v.port)
		}
	default:
		return nil, fmt.Errorf("Got: %T", x)
	}

	// (auth_checkPhone)
	resp = make(chan TL, 1)
	m.queueSend <- packetToSend{&TL_auth_checkPhone{"79197252476"}, resp}
	x = <-resp
	switch x.(type) {
	case *TL_rpc_error:
		x := x.(*TL_rpc_error)
		if x.error_code != 303 {
			return nil, fmt.Errorf("RPC error_code: %d", x.error_code)
		}
		var newDc int32
		n, _ := fmt.Sscanf(x.error_message, "NETWORK_MIGRATE_%d", &newDc)
		if n != 1 {
			return nil, fmt.Errorf("RPC error_string: %s", x.error_message)
		}
		newDcAddr, ok := m.dclist[newDc]
		if !ok {
			return nil, fmt.Errorf("Wrong DC index: %s", newDc)
		}
		// reconnect to newDcAddr
		fmt.Println(newDcAddr)
	default:
		return nil, fmt.Errorf("Got: %T", x)
	}

	// goroutine (TL_ping)
	go func() {
		for {
			select {
			case <-time.After(60 * time.Second):
				m.queueSend <- packetToSend{&TL_ping{0xCADACADA}, nil}
			}
		}
	}()

	return m, nil
}

func (m *MTProto) SendRoutine() {
	for x := range m.queueSend {
		needAck := true
		switch x.msg.(type) {
		case *TL_ping, *TL_msgs_ack:
			needAck = false
		}
		err := m.SendPacket(x.msg, needAck, x.resp)
		if err != nil {
			fmt.Println("SendRoutine:", err)
			os.Exit(2)
		}
	}
}

func (m *MTProto) ReadRoutine() {
	for true {
		data, err := m.Read()
		if err != nil {
			fmt.Println("ReadRoutine:", err)
			os.Exit(2)
		}

		m.Process(m.msgId, m.seqNo, data)
	}

}

func (m *MTProto) Process(msgId int64, seqNo int32, data interface{}) interface{} {
	switch data.(type) {
	case *TL_msg_container:
		data := data.(*TL_msg_container).items
		for _, v := range data {
			m.Process(v.msg_id, v.seq_no, v.data)
		}

	case *TL_bad_server_salt:
		data := data.(*TL_bad_server_salt)
		m.setSalt(data.new_server_salt)
		// TODO: resend messages

	case *TL_new_session_created:
		data := data.(*TL_new_session_created)
		m.setSalt(data.server_salt)

	case *TL_ping:
		data := data.(*TL_ping)
		m.queueSend <- packetToSend{&TL_pong{msgId, data.ping_id}, nil}

	case *TL_pong:
		// (ignore)

	case *TL_msgs_ack:
		data := data.(*TL_msgs_ack)
		for _, v := range data.msgIds {
			delete(m.msgsIdToAck, v)
		}

	case *TL_rpc_result:
		data := data.(*TL_rpc_result)
		x := m.Process(msgId, seqNo, data.obj)
		v, ok := m.msgsIdToResp[data.req_msg_id]
		if ok {
			v <- x.(TL)
			close(v)
			delete(m.msgsIdToResp, data.req_msg_id)
		}
		delete(m.msgsIdToAck, data.req_msg_id)

	default:
		return data

	}

	if (seqNo & 1) == 1 {
		m.SendPacket(&TL_msgs_ack{[]int64{msgId}}, false, nil)
	}

	return nil
}

func (m *MTProto) setGAB(g_ab *big.Int) {
	m.authKey = g_ab.Bytes()
	if m.authKey[0] == 0 {
		m.authKey = m.authKey[1:]
	}
	m.authKeyHash = sha1(m.authKey)[12:20]
	m.encrypted = g_ab.Cmp(big.NewInt(0)) != 0
	m.f.WriteAt(m.authKey, 0)
	m.f.WriteAt(m.authKeyHash, 256)
}

func (m *MTProto) setSalt(s []byte) {
	m.serverSalt = s
	m.f.WriteAt(m.serverSalt, 256+8)
}

func (m *MTProto) Halt() {
	select {}
}
