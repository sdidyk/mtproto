package mtproto

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"runtime"
	"sync"
	"time"
	"reflect"
	"github.com/k0kubun/pp"
)

const (
	appId   = 45139
	appHash = "7e55cea996fe1d94d6d22105258e3579"
	defaultServerAddr = "149.154.167.50:443"
)

type MTProto struct {
	addr      string
	conn      *net.TCPConn
	f         *os.File
	connected bool
	queueSend chan packetToSend
	stopRead  chan struct{}
	stopPing  chan struct{}

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
		m.addr = defaultServerAddr
		m.encrypted = false
	}
	rand.Seed(time.Now().UnixNano())
	m.sessionId = rand.Int63()
	m.connected = false
	return m, nil
}

func (m *MTProto) Connect() error {
	var err error
	var tcpAddr *net.TCPAddr
	//fmt.Println("Connecting: ", m.addr)

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
	m.stopRead = make(chan struct{}, 1)
	m.stopPing = make(chan struct{}, 1)
	m.msgsIdToAck = make(map[int64]packetToSend)
	m.msgsIdToResp = make(map[int64]chan TL)
	m.mutex = &sync.Mutex{}
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
				"Unknown",
				runtime.GOOS + "/" + runtime.GOARCH,
				"0.0.3",
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
	m.startPing()
	m.connected = true
	return nil
}

func (m *MTProto) Reconnect(newaddr string) error {
	var err error
	//fmt.Println("Reconnecting: ", newaddr)
	// stop ping routine
	m.stopPing <- struct{}{}
	close(m.stopPing)

	// close send routine & close connection
	close(m.queueSend)
	err = m.conn.Close()
	if err != nil {
		return err
	}
	m.connected = false
	// stop read routine
	m.stopRead <- struct{}{}
	close(m.stopRead)

	// renew connection
	m.encrypted = false
	m.addr = newaddr
	err = m.Connect()
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
				return fmt.Errorf("RPC error_code: %d, %s", x.error_code, x.error_message)
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
			err := m.Reconnect(newDcAddr)
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

		return fmt.Errorf("Cannot sign up yet")
	}

	return nil
}

func (m *MTProto) GetDialogs() error {
	resp := make(chan TL, 1)
	m.queueSend <- packetToSend{TL_messages_getDialogs{}, resp}
	x := <-resp

	list, ok := x.(TL_messages_dialogsSlice)
	if !ok {
		return fmt.Errorf("RPC: %#v", ok)
	}

	//fmt.Printf("%#v",list.chats)

	for _, chat := range list.chats {
		chat := chat.(TL_chat)
		fmt.Printf("%#v\n",chat.title)
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
		switch v.(type) {
	        case TL_userContact:
                v := v.(TL_userContact)
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

func (m *MTProto) GetFullChat(chat_id int32) error {
	resp := make(chan TL, 1)
	m.queueSend <- packetToSend{
		TL_messages_getFullChat{
			chat_id,
		}, 
		resp,
	}
	x := <-resp
	list, ok := x.(TL_messages_chatFull)
	if !ok {
		return fmt.Errorf("RPC: %#v", x)
	}

	fmt.Printf("%#v",list)

	return nil
}
func (m *MTProto) GetState() error {
	resp := make(chan TL, 1)
	m.queueSend <- packetToSend{TL_updates_getState{}, resp}
	x := <-resp
	_, ok := x.(TL_updates_state)
	if !ok {
		return fmt.Errorf("RPC: %#v", x)
	}

	//fmt.Printf("%#v",list)

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
		err := m.SendPacket(x.msg, x.resp)
		if err != nil {
			fmt.Fprintln(os.Stderr, "SendRoutine:", err)
			os.Exit(2)
		}
	}
}

func (m *MTProto) ReadRoutine(stop <-chan struct{}) {
	for true {
		data, err := m.Read(stop)
		if err != nil {
			fmt.Fprintln(os.Stderr, "ReadRoutine:", err)
			os.Exit(2)
		}
		if data == nil {
			return
		}

		m.Process(m.msgId, m.seqNo, data)
	}

}

func (m *MTProto) Process(msgId int64, seqNo int32, data interface{}) interface{} {
	fmt.Fprintln(os.Stderr, "Received: ", reflect.TypeOf(data))
	switch data.(type) {
	case TL_msg_container:
		data := data.(TL_msg_container).items
		for _, v := range data {
			m.Process(v.msg_id, v.seq_no, v.data)
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
		x := m.Process(msgId, seqNo, data.obj)
		m.mutex.Lock()
		v, ok := m.msgsIdToResp[data.req_msg_id]
		if ok {
			v <- x.(TL)
			close(v)
			delete(m.msgsIdToResp, data.req_msg_id)
		}
		delete(m.msgsIdToAck, data.req_msg_id)
		m.mutex.Unlock()
	
	case TL_rpc_error:
		if data == nil {
			break
		}
		data := data.(TL_rpc_error)
		fmt.Fprintln(os.Stderr, "RPC error: ", data)
	default:
		return data

	}

	if (seqNo & 1) == 1 && m.connected {
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
		return fmt.Errorf("New session")
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

func (m *MTProto) Halt() {
	select {}
}

func dump(x interface{}) {
	_, _ = pp.Println(x)
}
