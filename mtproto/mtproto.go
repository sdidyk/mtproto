package mtproto

import (
	"fmt"
	"github.com/sdidyk/pp"
	"math/big"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

const (
	appId   = 2899
	appHash = "36722c72256a24c1225de00eb6a1ca74"
)

type MTProto struct {
	addr      string
	conn      *net.TCPConn
	f         *os.File
	queueSend chan packetToSend
	stopRead  chan struct{}
	stopPing  chan struct{}

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
	msg  TL
	resp chan TL
}

func NewMTProto(authkeyfile string) (*MTProto, error) {
	var err error
	m := new(MTProto)

	// try to read [authKey, serverSalt, dcAddr]
	m.f, err = os.OpenFile(authkeyfile, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}
	b := make([]byte, 256+8+8+32)
	n, err := m.f.Read(b)
	if n == 256+8+8+32 {
		m.authKey = b[:256]
		m.authKeyHash = b[256 : 256+8]
		m.serverSalt = b[256+8 : 256+8+8]
		m.addr = strings.TrimRight(string(b[256+8+8:]), "\x00")
		m.encrypted = true
	} else {
		m.addr = "149.154.175.50:443"
		m.encrypted = false
	}
	rand.Seed(time.Now().UnixNano())
	m.sessionId = rand.Int63()

	return m, nil
}

func (m *MTProto) Connect() error {
	var err error
	var tcpAddr *net.TCPAddr

	// connect
	fmt.Printf("Connecting to %s\n", m.addr)
	tcpAddr, err = net.ResolveTCPAddr("tcp", m.addr)
	if err != nil {
		return err
	}
	m.conn, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return err
	}
	_, err = m.conn.Write([]byte{0xef})
	if err != nil {
		return err
	}

	// get new authKey if need
	if !m.encrypted {
		err = m.makeAuthKey()
		if err != nil {
			return err
		}
		m.setAddr(m.addr)
	}

	// start goroutines
	m.queueSend = make(chan packetToSend, 64)
	m.stopRead = make(chan struct{}, 1)
	m.stopPing = make(chan struct{}, 1)
	m.msgsIdToAck = make(map[int64]bool)
	m.msgsIdToResp = make(map[int64]chan TL)
	go m.SendRoutine()
	go m.ReadRoutine(m.stopRead)

	var resp chan TL
	var x TL

	// (help_getConfig)
	resp = make(chan TL, 1)
	m.queueSend <- packetToSend{
		TL_invokeWithLayer{
			layer,
			TL_initConnection{
				appId,
				"MacBook Pro (Retina, Mid 2012)",
				"OS X 10.10.3",
				"0.0.1",
				"en",
				TL_help_getConfig{},
			},
		},
		resp,
	}
	x = <-resp
	switch x.(type) {
	case TL_config:
		m.dclist = make(map[int32]string, 5)
		for _, v := range x.(TL_config).dc_options {
			v := v.(TL_dcOption)
			m.dclist[v.id] = fmt.Sprintf("%s:%d", v.ip_address, v.port)
		}
	default:
		return fmt.Errorf("Got: %T", x)
	}

	m.startPing()

	return nil
}

func (m *MTProto) Reconnect(newaddr string) error {
	// stop ping routine
	m.stopPing <- struct{}{}
	close(m.stopPing)

	// close send routine & close connection
	close(m.queueSend)
	m.conn.Close()

	// stop read routine
	m.stopRead <- struct{}{}
	close(m.stopRead)

	// renew connection
	m.encrypted = false
	m.addr = newaddr
	err := m.Connect()
	return err
}

func (m *MTProto) Auth(phonenumber string) error {
	var authSentCode TL_auth_sentCode

	// (TL_auth_sendCode)
	flag := true
	for flag {
		resp := make(chan TL, 1)
		m.queueSend <- packetToSend{TL_auth_sendCode{phonenumber, 0, appId, appHash, "en"}, resp}
		x := <-resp
		switch x.(type) {
		case TL_auth_sentCode:
			authSentCode = x.(TL_auth_sentCode)
			flag = false
		case TL_rpc_error:
			x := x.(TL_rpc_error)
			if x.error_code != 303 {
				return fmt.Errorf("RPC error_code: %d", x.error_code)
			}
			var newDc int32
			n, _ := fmt.Sscanf(x.error_message, "PHONE_MIGRATE_%d", &newDc)
			if n != 1 {
				n, _ := fmt.Sscanf(x.error_message, "NETWORK_MIGRATE_%d", &newDc)
				if n != 1 {
					return fmt.Errorf("RPC error_string: %s", x.error_message)
				}
			}

			newDcAddr, ok := m.dclist[newDc]
			if !ok {
				return fmt.Errorf("Wrong DC index: %s", newDc)
			}
			m.Reconnect(newDcAddr)
		default:
			return fmt.Errorf("Got: %T", x)
		}

	}

	var code int

	fmt.Print("Enter code: ")
	fmt.Scanf("%d", &code)

	if toBool(authSentCode.phone_registered) {
		resp := make(chan TL, 1)
		m.queueSend <- packetToSend{
			TL_auth_signIn{phonenumber, authSentCode.phone_code_hash, fmt.Sprintf("%d", code)},
			resp,
		}
		x := <-resp
		auth, ok := x.(TL_auth_authorization)
		if !ok {
			return fmt.Errorf("RPC: %#v", x)
		}
		userSelf := auth.user.(TL_userSelf)
		fmt.Printf("Signed in: id %d name <%s %s>\n", userSelf.id, userSelf.first_name, userSelf.last_name)

	} else {

		return fmt.Errorf("Cannot sign up yet")
	}

	return nil
}

func (m *MTProto) GetContacts() error {
	resp := make(chan TL, 1)
	m.queueSend <- packetToSend{TL_contacts_getContacts{""}, resp}
	x := <-resp
	list, ok := x.(TL_contacts_contacts)
	if !ok {
		return fmt.Errorf("RPC: %#v", x)
	}

	contacts := make(map[int32]TL_userContact)
	for _, v := range list.users {
		v := v.(TL_userContact)
		contacts[v.id] = v
	}
	fmt.Printf(
		"%10s    %10s    %-30s    %-20s\n",
		"(id)", "(mutual)", "(name)", "(username)",
	)
	for _, v := range list.contacts {
		v := v.(TL_contact)
		fmt.Printf(
			"%10d    %10t    %-30s    %-20s\n",
			v.user_id,
			toBool(v.mutual),
			fmt.Sprintf("%s %s", contacts[v.user_id].first_name, contacts[v.user_id].last_name),
			contacts[v.user_id].username,
		)
	}

	return nil
}

func (m *MTProto) SendMsg(user_id int32, msg string) error {
	resp := make(chan TL, 1)
	m.queueSend <- packetToSend{
		TL_messages_sendMessage{
			// TL_inputPeerSelf{},
			TL_inputPeerContact{user_id},
			msg,
			rand.Int63(),
		},
		resp,
	}
	x := <-resp
	_, ok := x.(TL_messages_sentMessage)
	if !ok {
		return fmt.Errorf("RPC: %#v", x)
	}

	return nil
}

func (m *MTProto) startPing() {
	// goroutine (TL_ping)
	go func() {
		for {
			select {
			case <-m.stopPing:
				return
			case <-time.After(60 * time.Second):
				m.queueSend <- packetToSend{TL_ping{0xCADACADA}, nil}
			}
		}
	}()
}

func (m *MTProto) SendRoutine() {
	for x := range m.queueSend {
		needAck := true
		switch x.msg.(type) {
		case TL_ping, TL_msgs_ack:
			needAck = false
		}
		err := m.SendPacket(x.msg, needAck, x.resp)
		if err != nil {
			fmt.Println("SendRoutine:", err)
			os.Exit(2)
		}
	}
}

func (m *MTProto) ReadRoutine(stop <-chan struct{}) {
	for true {
		data, err := m.Read(stop)
		if err != nil {
			fmt.Println("ReadRoutine:", err)
			os.Exit(2)
		}
		if data == nil {
			return
		}

		m.Process(m.msgId, m.seqNo, data)
	}

}

func (m *MTProto) Process(msgId int64, seqNo int32, data interface{}) interface{} {
	switch data.(type) {
	case TL_msg_container:
		data := data.(TL_msg_container).items
		for _, v := range data {
			m.Process(v.msg_id, v.seq_no, v.data)
		}

	case TL_bad_server_salt:
		data := data.(TL_bad_server_salt)
		m.setSalt(data.new_server_salt)
		// TODO: resend messages

	case TL_new_session_created:
		data := data.(TL_new_session_created)
		m.setSalt(data.server_salt)

	case TL_ping:
		data := data.(TL_ping)
		m.queueSend <- packetToSend{TL_pong{msgId, data.ping_id}, nil}

	case TL_pong:
		// (ignore)

	case TL_msgs_ack:
		data := data.(TL_msgs_ack)
		for _, v := range data.msgIds {
			delete(m.msgsIdToAck, v)
		}

	case TL_rpc_result:
		data := data.(TL_rpc_result)
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
		m.SendPacket(TL_msgs_ack{[]int64{msgId}}, false, nil)
	}

	return nil
}

func (m *MTProto) setGAB(g_ab *big.Int) {
	m.encrypted = true
	m.authKey = g_ab.Bytes()
	if m.authKey[0] == 0 {
		m.authKey = m.authKey[1:]
	}
	m.authKeyHash = sha1(m.authKey)[12:20]
	m.f.WriteAt(m.authKey, 0)
	m.f.WriteAt(m.authKeyHash, 256)
}

func (m *MTProto) setSalt(s []byte) {
	m.serverSalt = s
	m.f.WriteAt(m.serverSalt, 256+8)
}

func (m *MTProto) setAddr(s string) {
	m.addr = s
	b := make([]byte, 32)
	copy(b, []byte(s))
	m.f.WriteAt(b, 256+8+8)
}

func (m *MTProto) Halt() {
	select {}
}

func dump(x interface{}) {
	pp.Println(x)
}
