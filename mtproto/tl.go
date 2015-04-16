package mtproto

import (
	"fmt"
	"math/big"
)

const (
	// системные конструкторы
	crc_bool_false           = 0xbc799737
	crc_bool_true            = 0x997275b5
	crc_vector               = 0x1cb5c415
	crc_msg_container        = 0x73f1f8dc
	crc_new_session_created  = 0x9ec20908
	crc_msgs_ack             = 0x62d6b459
	crc_rpc_result           = 0xf35c6d01
	crc_rpc_error            = 0x2144ca19
	crc_bad_msg_notification = 0xa7eff811
	crc_bad_server_salt      = 0xedab447b

	// конструкторы авторизации
	crc_req_pq                = 0x60469778
	crc_resPQ                 = 0x05162463
	crc_p_q_inner_data        = 0x83c95aec
	crc_req_DH_params         = 0xd712e4be
	crc_server_DH_params_ok   = 0xd0e8075c
	crc_server_DH_params_fail = 0x79cb045d
	crc_server_DH_inner_data  = 0xb5890dba
	crc_client_DH_inner_data  = 0x6643b654
	crc_set_client_DH_params  = 0xf5045f1f
	crc_dh_gen_ok             = 0x3bcbf734
	crc_dh_gen_retry          = 0x46dc1fb9
	crc_dh_gen_fail           = 0xa69dae02

	// пинги и понги
	crc_ping = 0x7abe77ec
	crc_pong = 0x347773c5

	// коды help
	crc_help_getConfig = 0xc4f9186b
	crc_config         = 0x232d5905
	crc_dcOption       = 0x2ec2a43c
)

type TL_message struct {
	msg_id int64
	seq_no int32
	size   int32
	data   interface{}
}

type TL_resPQ struct {
	nonce        []byte
	server_nonce []byte
	pq           *big.Int
	fingerprints []int64
}

type TL_server_DH_params_ok struct {
	nonce            []byte
	server_nonce     []byte
	encrypted_answer []byte
}

type TL_server_DH_inner_data struct {
	nonce        []byte
	server_nonce []byte
	g            int32
	dh_prime     *big.Int
	g_a          *big.Int
	server_time  int32
}

type TL_new_session_created struct {
	first_msg_id int64
	unique_id    int64
	server_salt  []byte
}

func (m *DecodeBuf) DecodeRecursive(level int) (r interface{}) {
	constructor := m.DecodeUInt()
	if m.err != nil {
		return nil
	}

	level++

	switch constructor {

	case crc_resPQ:
		nonce := m.DecodeBytes(16)
		server_nonce := m.DecodeBytes(16)
		pq := m.DecodeBigInt()
		fingerprints := m.DecodeVectorLong()
		r = &TL_resPQ{nonce, server_nonce, pq, fingerprints}
		if m.err != nil {
			return nil
		}

	case crc_server_DH_params_ok:
		nonce := m.DecodeBytes(16)
		server_nonce := m.DecodeBytes(16)
		encrypted_answer := m.DecodeStringBytes()
		r = &TL_server_DH_params_ok{nonce, server_nonce, encrypted_answer}
		if m.err != nil {
			return nil
		}

	case crc_server_DH_inner_data:
		nonce := m.DecodeBytes(16)
		server_nonce := m.DecodeBytes(16)
		g := m.DecodeInt()
		dh_prime := m.DecodeBigInt()
		g_a := m.DecodeBigInt()
		server_time := m.DecodeInt()
		r = &TL_server_DH_inner_data{nonce, server_nonce, g, dh_prime, g_a, server_time}
		if m.err != nil {
			return nil
		}

	case crc_dh_gen_ok:
		nonce := m.DecodeBytes(16)
		server_nonce := m.DecodeBytes(16)
		new_nonce_hash1 := m.DecodeBytes(16)
		r = &TL_dh_gen_ok{nonce, server_nonce, new_nonce_hash1}
		if m.err != nil {
			return nil
		}

	case crc_ping:
		ping_id := m.DecodeLong()
		r = &TL_ping{ping_id}
		if m.err != nil {
			return nil
		}

	case crc_pong:
		msg_id := m.DecodeLong()
		ping_id := m.DecodeLong()
		r = &TL_pong{msg_id, ping_id}
		if m.err != nil {
			return nil
		}

	case crc_msg_container:
		size := m.DecodeInt()
		arr := make([]TL_message, size)
		for i := int32(0); i < size; i++ {
			msg_id := m.DecodeLong()
			seq_no := m.DecodeInt()
			size := m.DecodeInt()
			data := m.DecodeRecursive(level)
			arr[i] = TL_message{msg_id, seq_no, size, data}
			if m.err != nil {
				return nil
			}
		}
		r = &arr

	case crc_new_session_created:
		msg_id := m.DecodeLong()
		uniq_id := m.DecodeLong()
		server_salt := m.DecodeBytes(8)
		r = &TL_new_session_created{msg_id, uniq_id, server_salt}
		if m.err != nil {
			return nil
		}

	default:
		m.err = fmt.Errorf("Неизвестный конструктор: %08x", constructor)
		return nil

	}

	return
}

type TL_dh_gen_ok struct {
	nonce           []byte
	server_nonce    []byte
	new_nonce_hash1 []byte
}

type TL_ping struct {
	ping_id int64
}

type TL_pong struct {
	msg_id  int64
	ping_id int64
}
