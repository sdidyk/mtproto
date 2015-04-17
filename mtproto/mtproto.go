package mtproto

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
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
	// соединение
	conn *net.TCPConn
	f    *os.File

	QueueSend chan PacketToSend

	// данные соединения
	g_ab        *big.Int
	authKey     []byte
	authKeyHash []byte
	serverSalt  []byte
	encrypted   bool
	sessionId   int64

	// (туда)
	lastSeqNo   int32
	msgsIdToAck map[int64]bool

	// (сюда)
	seqNo int32
	msgId int64

	// разобранная структура
	data interface{}
}

type PacketToSend struct {
	Msg     []byte
	NeedAck bool
}

func NewMTProto(addr, authkeyfile string) (*MTProto, error) {
	var err error
	var tcpAddr *net.TCPAddr

	m := new(MTProto)

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
	m.g_ab = big.NewInt(0)
	rand.Seed(time.Now().UnixNano())
	m.sessionId = rand.Int63()

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

	if !m.encrypted {
		err = m.makeAuthKey()
		if err != nil {
			return nil, err
		}
	}

	m.QueueSend = make(chan PacketToSend, 64)
	m.msgsIdToAck = make(map[int64]bool)

	go func() {
		for {
			select {
			case <-time.After(60 * time.Second):
				m.QueueSend <- PacketToSend{Encode_TL_ping(0xCADACADA), false}
			}
		}
	}()
	go m.SendRoutine()
	go m.ReadRoutine()

	return m, nil
}

func (m *MTProto) SendPacket(msg []byte, needAck bool) error {
	x := make([]byte, 0, 256)

	if m.encrypted {
		z := make([]byte, 0, 256)
		newMsgId := GenerateMessageId()
		z = append(z, m.serverSalt...)
		z = append(z, EncodeLong(m.sessionId)...)
		z = append(z, EncodeLong(newMsgId)...)
		if needAck {
			z = append(z, EncodeInt(m.lastSeqNo|1)...)
		} else {
			z = append(z, EncodeInt(m.lastSeqNo)...)
		}
		z = append(z, EncodeInt(int32(len(msg)))...)
		z = append(z, msg...)

		msgKey := Sha1(z)[4:20]
		aesKey, aesIV := generateAES(msgKey, m.authKey, false)

		y := make([]byte, len(z)+((16-(len(msg)%16))&15))
		copy(y, z)
		encryptedData, err := AES256IGE_encrypt(y, aesKey, aesIV)
		if err != nil {
			return err
		}

		m.lastSeqNo += 2
		if needAck {
			m.msgsIdToAck[newMsgId] = true
		}

		x = append(x, m.authKeyHash...)
		x = append(x, msgKey...)
		x = append(x, encryptedData...)

	} else {
		x = append(x, EncodeLong(0)...)
		x = append(x, EncodeLong(GenerateMessageId())...)
		x = append(x, EncodeInt(int32(len(msg)))...)
		x = append(x, msg...)

	}

	size := len(x) / 4
	if size < 127 {
		x = append([]byte{byte(size)}, x...)
	} else {
		x = append(EncodeInt(int32(size<<8|127)), x...)
	}
	_, err := m.conn.Write(x)
	if err != nil {
		return err
	}

	return nil
}

func (m *MTProto) makeAuthKey() error {
	var x []byte
	var err error
	var data interface{}

	// (send) req_pq
	nonceFirst := GenerateNonce(16)
	err = m.SendPacket(Encode_TL_req_pq(nonceFirst), false)
	if err != nil {
		return err
	}

	// (parse) resPQ
	data, err = m.Read()
	if err != nil {
		return err
	}
	res, ok := data.(*TL_resPQ)
	if !ok {
		return errors.New("Handshake: ожидался resPQ")
	}
	if !bytes.Equal(nonceFirst, res.nonce) {
		return errors.New("Handshake: не совпадает nonce")
	}
	found := false
	for _, b := range res.fingerprints {
		if uint64(b) == telegramPublicKey_FP {
			found = true
			break
		}
	}
	if !found {
		return errors.New("Handshake: нет отпечатка нужного ключа")
	}

	// (encoding) p_q_inner_data
	p, q := SplitPQ(res.pq)
	nonceSecond := GenerateNonce(32)
	nonceServer := res.server_nonce
	innerData1 := Encode_TL_p_q_inner_data(res.pq, p, q, nonceFirst, nonceServer, nonceSecond)

	x = make([]byte, 255)
	copy(x[0:], Sha1(innerData1))
	copy(x[20:], innerData1)
	encryptedData1 := RSAEncode(x)

	// (send) req_DH_params
	err = m.SendPacket(Encode_TL_req_DH_params(nonceFirst, nonceServer, p, q, telegramPublicKey_FP, encryptedData1), false)
	if err != nil {
		return err
	}

	// (parse) server_DH_params_{ok, fail}
	data, err = m.Read()
	if err != nil {
		return err
	}
	dh, ok := data.(*TL_server_DH_params_ok)
	if !ok {
		return errors.New("Handshake: ожидался server_DH_params_ok")
	}
	if !bytes.Equal(nonceFirst, dh.nonce) {
		return errors.New("Handshake: не совпадает nonce")
	}
	if !bytes.Equal(nonceServer, dh.server_nonce) {
		return errors.New("Handshake: не совпадает server_nonce")
	}
	t1 := make([]byte, 48)
	copy(t1[0:], nonceSecond)
	copy(t1[32:], nonceServer)
	hash1 := Sha1(t1)

	t2 := make([]byte, 48)
	copy(t2[0:], nonceServer)
	copy(t2[16:], nonceSecond)
	hash2 := Sha1(t2)

	t3 := make([]byte, 64)
	copy(t3[0:], nonceSecond)
	copy(t3[32:], nonceSecond)
	hash3 := Sha1(t3)

	tmpAESKey := make([]byte, 32)
	tmpAESIV := make([]byte, 32)

	copy(tmpAESKey[0:], hash1)
	copy(tmpAESKey[20:], hash2[0:12])

	copy(tmpAESIV[0:], hash2[12:20])
	copy(tmpAESIV[8:], hash3)
	copy(tmpAESIV[28:], nonceSecond[0:4])

	// (parse-thru) server_DH_inner_data
	decodedData, err := AES256IGE_decrypt(dh.encrypted_answer, tmpAESKey, tmpAESIV)
	if err != nil {
		return err
	}
	innerbuf := NewDecodeBuf(decodedData[20:])
	data = innerbuf.DecodeObject(0)
	if innerbuf.err != nil {
		return innerbuf.err
	}
	dhi, ok := data.(*TL_server_DH_inner_data)
	if !ok {
		return errors.New("Handshake: ожидался server_DH_inner_data")
	}
	if !bytes.Equal(nonceFirst, dhi.nonce) {
		return errors.New("Handshake: не совпадает nonce")
	}
	if !bytes.Equal(nonceServer, dhi.server_nonce) {
		return errors.New("Handshake: не совпадает server_nonce")
	}

	_, g_b, g_ab := MakeGAB(dhi.g, dhi.g_a, dhi.dh_prime)
	authKey := g_ab.Bytes()
	if authKey[0] == 0 {
		authKey = authKey[1:]
	}
	t4 := make([]byte, 32+1+8)
	copy(t4[0:], nonceSecond)
	t4[32] = 1
	copy(t4[33:], Sha1(authKey)[0:8])
	nonceHash1 := Sha1(t4)[4:20]
	serverSalt := make([]byte, 8)
	copy(serverSalt, nonceSecond[:8])
	Xor(serverSalt, nonceServer[:8])

	// (encoding) client_DH_inner_data
	innerData2 := Encode_TL_client_DH_inner_data(nonceFirst, nonceServer, 0, g_b)
	x = make([]byte, 20+len(innerData2)+(16-((20+len(innerData2))%16))&15)
	copy(x[0:], Sha1(innerData2))
	copy(x[20:], innerData2)
	encryptedData2, err := AES256IGE_encrypt(x, tmpAESKey, tmpAESIV)

	// (send) set_client_DH_params
	err = m.SendPacket(Encode_TL_set_client_DH_params(nonceFirst, nonceServer, encryptedData2), false)
	if err != nil {
		return err
	}

	// (parse) dh_gen_{ok, retry, fail}
	data, err = m.Read()
	if err != nil {
		return err
	}
	dhg, ok := data.(*TL_dh_gen_ok)
	if !ok {
		return errors.New("Handshake: ожидался dh_gen_ok")
	}
	if !bytes.Equal(nonceFirst, dhg.nonce) {
		return errors.New("Handshake: не совпадает nonce")
	}
	if !bytes.Equal(nonceServer, dhg.server_nonce) {
		return errors.New("Handshake: не совпадает server_nonce")
	}
	if !bytes.Equal(nonceHash1, dhg.new_nonce_hash1) {
		return errors.New("Handshake: не совпадает new_nonce_hash1")
	}

	// (all ok)
	m.setGAB(g_ab)
	m.setSalt(serverSalt)

	return nil
}

func (m *MTProto) Read() (interface{}, error) {
	var err error
	var n int
	var size int
	var data interface{}

	m.conn.SetReadDeadline(time.Now().Add(300 * time.Second))
	b := make([]byte, 1)
	n, err = m.conn.Read(b)
	if err != nil {
		return nil, err
	}

	if b[0] < 127 {
		size = int(b[0]) << 2
	} else {
		b := make([]byte, 3)
		n, err = m.conn.Read(b)
		if err != nil {
			return nil, err
		}
		size = (int(b[0]) | int(b[1])<<8 | int(b[2])<<16) << 2
	}

	left := size
	buf := make([]byte, size)
	for left > 0 {
		n, err = m.conn.Read(buf[size-left:])
		if err != nil {
			return nil, err
		}
		left -= n
	}

	if size == 4 {
		return nil, fmt.Errorf("Ошибка сервера: %d", int32(binary.LittleEndian.Uint32(buf)))
	}

	dbuf := NewDecodeBuf(buf)

	authKeyHash := dbuf.DecodeBytes(8)
	if binary.LittleEndian.Uint64(authKeyHash) == 0 {
		m.msgId = dbuf.DecodeLong()
		messageLen := dbuf.DecodeInt()
		if int(messageLen) != dbuf.size-20 {
			return nil, fmt.Errorf("Длина сообщения не совпадает: %d (должна быть %d)", messageLen, dbuf.size-20)
		}
		m.seqNo = 0

		data = dbuf.DecodeObject(0)
		if dbuf.err != nil {
			return nil, dbuf.err
		}

	} else {
		msgKey := dbuf.DecodeBytes(16)
		encryptedData := dbuf.DecodeBytes(dbuf.size - 24)
		aesKey, aesIV := generateAES(msgKey, m.authKey, true)
		x, err := AES256IGE_decrypt(encryptedData, aesKey, aesIV)
		if err != nil {
			return nil, err
		}
		dbuf = NewDecodeBuf(x)
		_ = dbuf.DecodeLong() // salt
		_ = dbuf.DecodeLong() // session_id
		m.msgId = dbuf.DecodeLong()
		m.seqNo = dbuf.DecodeInt()
		messageLen := dbuf.DecodeInt()
		if int(messageLen) >= +dbuf.size-32 {
			return nil, fmt.Errorf("Длина сообщения не совпадает: %d (максимум %d)", messageLen, dbuf.size-32)
		}
		if !bytes.Equal(Sha1(dbuf.buf[0 : 32+messageLen])[4:20], msgKey) {
			return nil, fmt.Errorf("Битый msg_key")
		}

		data = dbuf.DecodeObject(0)
		if dbuf.err != nil {
			return nil, dbuf.err
		}

	}
	mod := m.msgId & 3
	if mod != 1 && mod != 3 {
		return nil, fmt.Errorf("Невалидные битые message_id: %d", mod)
	}

	return data, nil
}

func (m *MTProto) SendRoutine() {
	for x := range m.QueueSend {
		m.SendPacket(x.Msg, x.NeedAck)
	}
}

func (m *MTProto) ReadRoutine() {
	for true {
		data, err := m.Read()
		if err != nil {
			fmt.Println("ReadRoutine:", err)
			os.Exit(2)
		}

		switch data.(type) {

		case []TL_message:
			data := data.([]TL_message)
			for _, v := range data {
				m.Process(v.msg_id, v.seq_no, v.data)
			}

		default:
			m.Process(m.msgId, m.seqNo, data)

		}
	}

}

func (m *MTProto) Process(msgId int64, seqNo int32, data interface{}) {
	switch data.(type) {

	case *TL_bad_server_salt:
		data := data.(*TL_bad_server_salt)
		m.setSalt(data.new_server_salt)

	case *TL_new_session_created:
		data := data.(*TL_new_session_created)
		m.setSalt(data.server_salt)

	case *TL_ping:
		data := data.(*TL_ping)
		m.QueueSend <- PacketToSend{Encode_TL_pong(msgId, data.ping_id), false}

	case *TL_pong:
		// ignore

	case *TL_msgs_ack:
		data := data.(*TL_msgs_ack)
		for _, v := range data.msgIds {
			delete(m.msgsIdToAck, v)
		}

	case *TL_rpc_result:
		data := data.(*TL_rpc_result)
		delete(m.msgsIdToAck, data.req_msg_id)
		m.Process(msgId, seqNo, data.obj)

	default:
		fmt.Println("INFO: data to process", data)

	}

	if (seqNo & 1) == 1 {
		m.SendPacket(Encode_TL_msgs_ack([]int64{msgId}), false)
	}

}

func (m *MTProto) setGAB(g_ab *big.Int) {
	m.g_ab = g_ab
	m.authKey = g_ab.Bytes()
	if m.authKey[0] == 0 {
		m.authKey = m.authKey[1:]
	}
	m.authKeyHash = Sha1(m.authKey)[12:20]
	m.encrypted = g_ab.Cmp(big.NewInt(0)) != 0
	m.f.WriteAt(m.authKey, 0)
	m.f.WriteAt(m.authKeyHash, 256)
}

func (m *MTProto) setSalt(s []byte) {
	m.serverSalt = s
	m.f.WriteAt(m.serverSalt, 256+8)
}

func Dump(x []byte) {
	fmt.Println(hex.Dump(x))
}
