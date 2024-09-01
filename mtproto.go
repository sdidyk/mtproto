package mtproto

import (
	"errors"
	"fmt"
	"math/rand"
	"net"
	"os"
	"runtime"
	"slices"
	"sync"
	"time"
)

const (
	appId   = 41994
	appHash = "269069e15c81241f5670c397941016a2"
)

type MTProto struct {
	addr      string
	dc        int32
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
		m.dc = 2
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
		err = m.makeAuthKey(m.dc)
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
				api_id:           appId,
				device_model:     "Unknown",
				system_version:   runtime.GOOS + "/" + runtime.GOARCH,
				app_version:      "0.1.0",
				system_lang_code: "en",
				lang_pack:        "en",
				lang_code:        "en",
				query:            TL_help_getConfig{},
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
			if !v.ipv6 && !v.media_only && !v.tcpo_only && !v.cdn && !v.static {
				m.dclist[v.id] = fmt.Sprintf("%s:%d", v.ip_address, v.port)
			}
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
	var x TL
	var resp chan TL

	resp = make(chan TL, 1)
	m.queueSend <- packetToSend{TL_auth_sendCode{phonenumber, appId, appHash, &TL_codeSettings{}}, resp}
	x = <-resp

	authSentCode, ok := x.(TL_auth_sentCode)
	if !ok {
		return fmt.Errorf("RPC: %#v", x)
	}
	fmt.Println(authSentCode)

	var code int

	fmt.Print("Enter code: ")
	fmt.Scanf("%d", &code)

	fmt.Println("Login:", phonenumber, authSentCode.phone_code_hash, code)
	resp = make(chan TL, 1)
	m.queueSend <- packetToSend{
		TL_auth_signIn{
			phone_number: phonenumber,
			phone_code_hash: authSentCode.phone_code_hash,
			phone_code: fmt.Sprintf("%d", code),
		},
		resp,
	}
	x = <-resp
	auth, ok := x.(TL_auth_authorization)
	if !ok {
		return fmt.Errorf("RPC: %#v", x)
	}
	user := auth.user.(TL_user)
	fmt.Printf("Signed in: id %d name <%s %s>\n", user.id, user.first_name, user.last_name)

	return nil
}

func (m *MTProto) GetContacts() error {
	resp := make(chan TL, 1)
	m.queueSend <- packetToSend{TL_contacts_getContacts{0}, resp}
	x := <-resp
	list, ok := x.(TL_contacts_contacts)
	if !ok {
		return fmt.Errorf("RPC: %#v", x)
	}

	users := make(map[int64]TL_user)
	for _, v := range list.users {
		if v, ok := v.(TL_user); ok {
			users[v.id] = v
		}
	}

	userKeys := make([]int64, 0, len(users))
	for k := range users {
		userKeys = append(userKeys, k)
	}
	slices.Sort(userKeys)

	fmt.Printf(
		"\033[33m\033[1m%10s    %8s    %-30s    %-30s    %16s\033[0m\n",
		"id", "mutual", "name", "username", "access_hash",
	)
	for _, uid := range userKeys {
		v := users[uid]
		fmt.Printf(
			"%10d    %8t    %-30s    %-30s    %016x\n",
			v.id,
			v.mutual_contact,
			fmt.Sprintf("%s %s", v.first_name, v.last_name),
			v.username,
			uint64(v.access_hash),
		)
	}

	return nil
}

func (m *MTProto) SendMessage(user_id int64, access_hash int64, msg string) error {
	resp := make(chan TL, 1)
	m.queueSend <- packetToSend{
		TL_messages_sendMessage{
			peer:      TL_inputPeerUser{user_id, access_hash},
			message:   msg,
			random_id: rand.Int63(),
		},
		resp,
	}
	x := <-resp
	fmt.Printf("RPC: %#v\n", x)

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
