package mtproto

import (
	"math/big"
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
	date               int32
	expires            int32
	test_mode          bool
	this_dc            int32
	dc_options         []TL_dcOption
	chat_big_size      int32
	chat_size_max      int32
	broadcast_size_max int32
	disabled_features  []TL_disabledFeature
}

type TL_dcOption struct {
	id         int32
	hostname   string
	ip_address string
	port       int32
}

type TL_disabledFeature struct {
	feature     string
	description string
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

type TL_invokeWithLayer struct {
	layer int32
	query TL
}

type TL_initConnection struct {
	app_id         int32
	device_model   string
	system_version string
	app_version    string
	lang_code      string
	query          TL
}
