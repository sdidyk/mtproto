package mtproto

import (
	"crypto/aes"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
)

const (
	// системные конструкторы
	bool_false           = 0xbc799737
	bool_true            = 0x997275b5
	vector               = 0x1cb5c415
	msg_container        = 0x73f1f8dc
	new_session_created  = 0x9ec20908
	msgs_ack             = 0x62d6b459
	rpc_result           = 0xf35c6d01
	rpc_error            = 0x2144ca19
	bad_msg_notification = 0xa7eff811
	bad_server_salt      = 0xedab447b

	// конструкторы авторизации
	req_pq                = 0x60469778
	resPQ                 = 0x05162463
	p_q_inner_data        = 0x83c95aec
	req_DH_params         = 0xd712e4be
	server_DH_params_ok   = 0xd0e8075c
	server_DH_params_fail = 0x79cb045d
	server_DH_inner_data  = 0xb5890dba
	client_DH_inner_data  = 0x6643b654
	set_client_DH_params  = 0xf5045f1f
	dh_gen_ok             = 0x3bcbf734
	dh_gen_retry          = 0x46dc1fb9
	dh_gen_fail           = 0xa69dae02
)

type TL_resPQ struct {
	nonce        []byte
	server_nonce []byte
	pq           *big.Int
	fingerprints []uint64
}

type TL_server_DH_params_ok struct {
	nonce            []byte
	server_nonce     []byte
	encrypted_answer []byte
}

func AES256IGE_decrypt(data, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(data) < aes.BlockSize {
		return nil, errors.New("Слишком короткие данные")
	}
	if len(data)%aes.BlockSize != 0 {
		return nil, errors.New("Данные некратны блоку")
	}

	t := make([]byte, aes.BlockSize)
	x := iv[:aes.BlockSize]
	y := iv[aes.BlockSize:]

	decrtypted := make([]byte, len(data))

	i := 0
	for i < len(data) {
		xor(y, data[i:i+aes.BlockSize])
		block.Decrypt(t, y)
		xor(t, x)
		y, x = t, data[i:i+aes.BlockSize]
		copy(decrtypted[i:], t)
		i += aes.BlockSize
	}

	return decrtypted, nil

}

func xor(dst, src []byte) {
	for i, _ := range dst {
		dst[i] = dst[i] ^ src[i]
	}
}

func (m *MTProto) DecodePacket() error {
	var err error

	constructor, err := m.DecodeUInt()
	if err != nil {
		return err
	}

	m.level++

	switch constructor {
	case resPQ:
		nonce, err := m.DecodeBytes(16)
		server_nonce, err := m.DecodeBytes(16)
		pq, err := m.DecodeBigInt()
		fingerprints, err := m.DecodeVectorLong()
		m.data = TL_resPQ{nonce, server_nonce, pq, fingerprints}
		if err != nil {
			return err
		}
	case server_DH_params_ok:
		nonce, err := m.DecodeBytes(16)
		server_nonce, err := m.DecodeBytes(16)
		encrypted_answer, err := m.DecodeStringBytes()
		m.data = TL_server_DH_params_ok{nonce, server_nonce, encrypted_answer}
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("Неизвестный конструктор: %08x", constructor)
	}

	m.level--

	return nil
}

func (m *MTProto) DecodeLong() (r uint64, err error) {
	if m.off+8 > m.size {
		return 0, errors.New("DecodeLong: короткий пакет")
	}
	x := binary.LittleEndian.Uint64(m.buf[m.off : m.off+8])
	m.off += 8
	return x, nil
}

func (m *MTProto) DecodeInt() (r int32, err error) {
	if m.off+4 > m.size {
		return 0, errors.New("DecodeInt: короткий пакет")
	}
	x := binary.LittleEndian.Uint32(m.buf[m.off : m.off+4])
	m.off += 4
	return int32(x), nil
}

func (m *MTProto) DecodeUInt() (r uint32, err error) {
	if m.off+4 > m.size {
		return 0, errors.New("DecodeUInt: короткий пакет")
	}
	x := binary.LittleEndian.Uint32(m.buf[m.off : m.off+4])
	m.off += 4
	return x, nil
}

func (m *MTProto) DecodeBytes(size int) (r []byte, err error) {
	if m.off+size > m.size {
		return nil, errors.New("DecodeBytes: короткий пакет")
	}
	x := make([]byte, size)
	copy(x, m.buf[m.off:m.off+size])
	m.off += size
	return x, nil
}

func (m *MTProto) DecodeStringBytes() (r []byte, err error) {
	var size, padding int

	if m.off+1 > m.size {
		return nil, errors.New("DecodeStringBytes: короткий пакет")
	}
	size = int(m.buf[m.off])
	m.off++
	padding = (4 - ((size + 1) % 4)) & 3
	if size == 254 {
		if m.off+3 > m.size {
			return nil, errors.New("DecodeStringBytes: короткий пакет")
		}
		size = int(m.buf[m.off]) | int(m.buf[m.off+1])<<8 | int(m.buf[m.off+2])<<16
		m.off += 3
		padding = (4 - size%4) & 3
	}

	if m.off+size > m.size {
		return nil, errors.New("DecodeStringBytes: короткий пакет (size)")
	}
	x := make([]byte, size)
	copy(x, m.buf[m.off:m.off+size])
	m.off += size

	if m.off+padding > m.size {
		return nil, errors.New("DecodeStringBytes: короткий пакет (padding)")
	}
	m.off += padding

	return x, nil
}

func (m *MTProto) DecodeBigInt() (r *big.Int, err error) {
	b, err := m.DecodeStringBytes()
	if err != nil {
		return nil, err
	}
	y := make([]byte, len(b)+1)
	y[0] = 0
	copy(y[1:], b)
	x := new(big.Int).SetBytes(y)
	return x, nil
}

func (m *MTProto) DecodeVectorLong() (r []uint64, err error) {
	constructor, err := m.DecodeUInt()
	if err != nil {
		return nil, err
	}
	if constructor != vector {
		return nil, errors.New("DecodeVectorLong: Неправильный конструктор")
	}
	size, err := m.DecodeInt()
	if err != nil {
		return nil, err
	}
	if size <= 0 {
		return nil, errors.New("DecodeVectorLong: Неправильный размер")
	}
	x := make([]uint64, size)
	i := int32(0)
	for i < size {
		y, err := m.DecodeLong()
		if err != nil {
			return nil, err
		}
		x[i] = y
		i++
	}
	return x, nil
}
