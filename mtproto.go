package mtproto

import (
	"errors"
	"fmt"
	"math/rand"
	"net"
	"os"
	"runtime"
	"sync"
	"time"
)

const (
	appId   = 41994
	appHash = "269069e15c81241f5670c397941016a2"
)

type MTProto struct {
	addr      string
	conn      *net.TCPConn
	f         *os.File
	queueSend chan packetToSend
	stopSend  chan struct{}
	stopRead  chan struct{}
	stopPing  chan struct{}
	allDone   chan struct{}

	authKey     []byte
	authKeyHash []byte
	serverSalt  []byte
	encrypted   bool
	sessionId   int64

	mutex        *sync.Mutex
	lastSeqNo    int32
	msgsIdToAck  map[int64]packetToSend
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

	m.f, err = os.OpenFile(authkeyfile, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}

	err = m.readData()
	if err == nil {
		m.encrypted = true
	} else {
		m.addr = "149.154.167.50:443"
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
	}

	// start goroutines
	m.queueSend = make(chan packetToSend, 64)
	m.stopSend = make(chan struct{}, 1)
	m.stopRead = make(chan struct{}, 1)
	m.stopPing = make(chan struct{}, 1)
	m.allDone = make(chan struct{}, 3)
	m.msgsIdToAck = make(map[int64]packetToSend)
	m.msgsIdToResp = make(map[int64]chan TL)
	m.mutex = &sync.Mutex{}
	go m.sendRoutine()
	go m.readRoutine()

	var resp chan TL
	var x TL

	// (help_getConfig)
	resp = make(chan TL, 1)
	m.queueSend <- packetToSend{
		TL_invokeWithLayer{
			layer,
			TL_initConnection{
				appId,
				"Unknown",
				runtime.GOOS + "/" + runtime.GOARCH,
				"0.0.4",
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

	// start keepalive pinging
	go m.pingRoutine()

	return nil
}

func (m *MTProto) reconnect(newaddr string) error {
	var err error

	// stop ping routine
	m.stopPing <- struct{}{}
	close(m.stopPing)

	// stop send routine
	m.stopSend <- struct{}{}
	close(m.stopSend)

	// stop read routine
	m.stopRead <- struct{}{}
	close(m.stopRead)

	<-m.allDone
	<-m.allDone
	<-m.allDone

	// close send queue
	close(m.queueSend)

	// close connection
	err = m.conn.Close()
	if err != nil {
		return err
	}

	// renew connection
	m.encrypted = false
	m.addr = newaddr
	err = m.Connect()
	return err
}

func (m *MTProto) Auth(phonenumber string) error {
	var authSentCode TL_auth_sentCode

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
				return fmt.Errorf("Wrong DC index: %d", newDc)
			}
			err := m.reconnect(newDcAddr)
			if err != nil {
				return err
			}
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

		return errors.New("Cannot sign up yet")
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
		if v, ok := v.(TL_userContact); ok {
			contacts[v.id] = v
		}
	}
	fmt.Printf(
		"\033[33m\033[1m%10s    %10s    %-30s    %-20s\033[0m\n",
		"id", "mutual", "name", "username",
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

func (m *MTProto) SendMessage(user_id int32, msg string) error {
	resp := make(chan TL, 1)
	m.queueSend <- packetToSend{
		TL_messages_sendMessage{
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

func (m *MTProto) pingRoutine() {
	for {
		select {
		case <-m.stopPing:
			m.allDone <- struct{}{}
			return
		case <-time.After(60 * time.Second):
			m.queueSend <- packetToSend{TL_ping{0xCADACADA}, nil}
		}
	}
}

func (m *MTProto) sendRoutine() {
	for x := range m.queueSend {
		err := m.sendPacket(x.msg, x.resp)
		if err != nil {
			fmt.Println("SendRoutine:", err)
			os.Exit(2)
		}
	}

	m.allDone <- struct{}{}
}

func (m *MTProto) readRoutine() {
	for {
		data, err := m.read(m.stopRead)
		if err != nil {
			fmt.Println("ReadRoutine:", err)
			os.Exit(2)
		}
		if data == nil {
			m.allDone <- struct{}{}
			return
		}

		m.process(m.msgId, m.seqNo, data)
	}

}

func (m *MTProto) process(msgId int64, seqNo int32, data interface{}) interface{} {
	switch data.(type) {
	case TL_msg_container:
		data := data.(TL_msg_container).items
		for _, v := range data {
			m.process(v.msg_id, v.seq_no, v.data)
		}

	case TL_bad_server_salt:
		data := data.(TL_bad_server_salt)
		m.serverSalt = data.new_server_salt
		_ = m.saveData()
		m.mutex.Lock()
		for k, v := range m.msgsIdToAck {
			delete(m.msgsIdToAck, k)
			m.queueSend <- v
		}
		m.mutex.Unlock()

	case TL_new_session_created:
		data := data.(TL_new_session_created)
		m.serverSalt = data.server_salt
		_ = m.saveData()

	case TL_ping:
		data := data.(TL_ping)
		m.queueSend <- packetToSend{TL_pong{msgId, data.ping_id}, nil}

	case TL_pong:
		// (ignore)

	case TL_msgs_ack:
		data := data.(TL_msgs_ack)
		m.mutex.Lock()
		for _, v := range data.msgIds {
			delete(m.msgsIdToAck, v)
		}
		m.mutex.Unlock()

	case TL_rpc_result:
		data := data.(TL_rpc_result)
		x := m.process(msgId, seqNo, data.obj)
		m.mutex.Lock()
		v, ok := m.msgsIdToResp[data.req_msg_id]
		if ok {
			v <- x.(TL)
			close(v)
			delete(m.msgsIdToResp, data.req_msg_id)
		}
		delete(m.msgsIdToAck, data.req_msg_id)
		m.mutex.Unlock()

	default:
		return data

	}

	if (seqNo & 1) == 1 {
		m.queueSend <- packetToSend{TL_msgs_ack{[]int64{msgId}}, nil}
	}

	return nil
}

func (m *MTProto) saveData() (err error) {
	m.encrypted = true

	b := NewEncodeBuf(1024)
	b.StringBytes(m.authKey)
	b.StringBytes(m.authKeyHash)
	b.StringBytes(m.serverSalt)
	b.String(m.addr)

	err = m.f.Truncate(0)
	if err != nil {
		return err
	}

	_, err = m.f.WriteAt(b.buf, 0)
	if err != nil {
		return err
	}

	return nil
}

func (m *MTProto) readData() (err error) {
	b := make([]byte, 1024*4)
	n, err := m.f.ReadAt(b, 0)
	if n <= 0 {
		return errors.New("New session")
	}

	d := NewDecodeBuf(b)
	m.authKey = d.StringBytes()
	m.authKeyHash = d.StringBytes()
	m.serverSalt = d.StringBytes()
	m.addr = d.String()

	if d.err != nil {
		return d.err
	}

	return nil
}
