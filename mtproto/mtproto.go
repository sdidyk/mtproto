package mtproto

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"net"
	"reflect"
)

type MTProto struct {
	g_ab       *big.Int
	serverSalt uint64
	conn       *net.TCPConn
	encrypted  bool

	buf  []byte
	size int
	off  int

	level     int
	messageId uint64
	seqNo     int32

	data interface{}
}

func (m *MTProto) Connect(addr string) error {
	var err error
	var tcpAddr *net.TCPAddr

	m.g_ab = new(big.Int)
	m.serverSalt = 0
	m.encrypted = false

	tcpAddr, err = net.ResolveTCPAddr("tcp", addr)
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

	return nil
}

func (m *MTProto) SendPacket(msg []byte) error {
	x := make([]byte, 0, 256)

	if m.encrypted {
		// TODO: encrypt packet

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

func (m *MTProto) Handshake() error {
	var x []byte
	var err error

	// (send) req_pq
	nonceFirst := GenerateNonce(16)
	err = m.SendPacket(Encode_TL_req_pq(nonceFirst))
	if err != nil {
		return err
	}

	// (parse) resPQ
	err = m.Read()
	if err != nil {
		return err
	}
	res, ok := m.data.(TL_resPQ)
	if !ok {
		return errors.New("Handshake: ожидался TL_resPQ")
	}
	if !bytes.Equal(nonceFirst, res.nonce) {
		return errors.New("Handshake: не совпадает nonce")
	}
	found := false
	for _, b := range res.fingerprints {
		if b == telegramPublicKey_FP {
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
	err = m.SendPacket(Encode_TL_req_DH_params(nonceFirst, nonceServer, p, q, telegramPublicKey_FP, encryptedData1))
	if err != nil {
		return err
	}

	// (parse) server_DH_params_{ok, fail}
	err = m.Read()
	if err != nil {
		return err
	}

	// (send) set_client_DH_params
	// (parse) dh_gen_{ok, retry, fail}

	return nil
}

func (m *MTProto) Read() error {
	var err error
	var n int

	b := make([]byte, 1)
	n, err = m.conn.Read(b)
	if err != nil {
		return err
	}

	if b[0] < 127 {
		m.size = int(b[0]) << 2
	} else {
		b := make([]byte, 3)
		n, err = m.conn.Read(b)
		if err != nil {
			return err
		}
		m.size = (int(b[0]) | int(b[1])<<8 | int(b[2])<<16) << 2
	}

	left := m.size
	m.buf = make([]byte, m.size)
	for left > 0 {
		n, err = m.conn.Read(m.buf[m.size-left:])
		if err != nil {
			return err
		}
		left -= n
	}
	m.off = 0

	if m.size == 4 {
		return fmt.Errorf("Ошибка: %d", int32(binary.LittleEndian.Uint32(m.buf)))
	}

	if m.size <= 8 {
		return fmt.Errorf("Слишком маленький пакет: %d байт", m.size)
	}

	authKeyHash, err := m.DecodeLong()
	if authKeyHash == 0 {
		m.messageId, err = m.DecodeLong()
		if err != nil {
			return err
		}
		messageLen, err := m.DecodeInt()
		if err != nil {
			return err
		}
		if int(messageLen) != m.size-20 {
			return fmt.Errorf("Длина сообщения не совпадает: %d (должна быть %d)", messageLen, m.size-20)
		}
		mod := m.messageId & 3
		if mod != 1 && mod != 3 {
			return fmt.Errorf("Невалидные битые message_id: %d", mod)
		}
		m.seqNo = 0
		m.level = 0

		err = m.DecodePacket()
		if err != nil {
			return err
		}

	} else {
		// TODO: read encrypted packet

	}

	return nil
}

func (m *MTProto) Dump() {
	fmt.Printf("AB: %v\tSALT: %v\tConnection: %v\n", m.g_ab, m.serverSalt, m.conn)
	fmt.Println(reflect.TypeOf(m.data), m.data)
}

func Dump(x []byte) {
	fmt.Println(hex.Dump(x))
}
