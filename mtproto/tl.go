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

func (m *MTProto) DecodePacket() error {
	var err error

	constructor, err := m.DecodeUInt()
	if err != nil {
		return err
	}

	m.level++

	switch constructor {
	case crc_resPQ:
		nonce, err := m.DecodeBytes(16)
		server_nonce, err := m.DecodeBytes(16)
		pq, err := m.DecodeBigInt()
		fingerprints, err := m.DecodeVectorLong()
		m.data = &TL_resPQ{nonce, server_nonce, pq, fingerprints}
		if err != nil {
			return err
		}

	case crc_server_DH_params_ok:
		nonce, err := m.DecodeBytes(16)
		server_nonce, err := m.DecodeBytes(16)
		encrypted_answer, err := m.DecodeStringBytes()
		m.data = &TL_server_DH_params_ok{nonce, server_nonce, encrypted_answer}
		if err != nil {
			return err
		}

	case crc_server_DH_inner_data:
		nonce, err := m.DecodeBytes(16)
		server_nonce, err := m.DecodeBytes(16)
		g, err := m.DecodeInt()
		dh_prime, err := m.DecodeBigInt()
		g_a, err := m.DecodeBigInt()
		server_time, err := m.DecodeInt()
		m.data = &TL_server_DH_inner_data{nonce, server_nonce, g, dh_prime, g_a, server_time}
		if err != nil {
			return err
		}

	case crc_dh_gen_ok:
		nonce, err := m.DecodeBytes(16)
		server_nonce, err := m.DecodeBytes(16)
		new_nonce_hash1, err := m.DecodeBytes(16)
		m.data = &TL_dh_gen_ok{nonce, server_nonce, new_nonce_hash1}
		if err != nil {
			return err
		}

	case crc_ping:
		ping_id, err := m.DecodeLong()
		m.data = &TL_ping{ping_id}
		if err != nil {
			return err
		}

	case crc_pong:
		msg_id, err := m.DecodeLong()
		ping_id, err := m.DecodeLong()
		m.data = &TL_pong{msg_id, ping_id}
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("Неизвестный конструктор: %08x", constructor)
	}

	m.level--

	return nil
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
