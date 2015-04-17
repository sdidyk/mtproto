package mtproto

import (
	"math/big"
)

const (
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

	crc_ping = 0x7abe77ec
	crc_pong = 0x347773c5

	crc_help_getConfig = 0xc4f9186b
	crc_config         = 0x232d5905
	crc_dcOption       = 0x2ec2a43c
)

type TL interface {
	encode() []byte
}

type TL_message struct {
	msg_id int64
	seq_no int32
	size   int32
	data   interface{}
}

type TL_req_pq struct {
	nonce []byte
}

type TL_p_q_inner_data struct {
	pq           *big.Int
	p            *big.Int
	q            *big.Int
	nonce        []byte
	server_nonce []byte
	new_nonce    []byte
}
type TL_req_DH_params struct {
	nonce        []byte
	server_nonce []byte
	p            *big.Int
	q            *big.Int
	fp           uint64
	encdata      []byte
}
type TL_client_DH_inner_data struct {
	nonce        []byte
	server_nonce []byte
	retry        int64
	g_b          *big.Int
}
type TL_set_client_DH_params struct {
	nonce        []byte
	server_nonce []byte
	encdata      []byte
}
type TL_help_getConfig struct{}

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

type TL_bad_server_salt struct {
	bad_msg_id      int64
	bad_msg_seqno   int32
	error_code      int32
	new_server_salt []byte
}

type TL_msgs_ack struct {
	msgIds []int64
}

type TL_rpc_result struct {
	req_msg_id int64
	obj        interface{}
}

type TL_config struct {
	date          int32
	test_mode     bool
	this_dc       int32
	dc_options    []TL_dcOption
	chat_size_max int32
}

type TL_dcOption struct {
	id         int32
	hostname   string
	ip_address string
	port       int32
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
