package mtproto

const (
	layer = 23

	// https://core.telegram.org/schema/mtproto
	crc_vector                     = 0x1cb5c415
	crc_resPQ                      = 0x05162463
	crc_p_q_inner_data             = 0x83c95aec
	crc_server_DH_params_fail      = 0x79cb045d
	crc_server_DH_params_ok        = 0xd0e8075c
	crc_server_DH_inner_data       = 0xb5890dba
	crc_client_DH_inner_data       = 0x6643b654
	crc_dh_gen_ok                  = 0x3bcbf734
	crc_dh_gen_retry               = 0x46dc1fb9
	crc_dh_gen_fail                = 0xa69dae02
	crc_rpc_result                 = 0xf35c6d01
	crc_rpc_error                  = 0x2144ca19
	crc_rpc_answer_unknown         = 0x5e2ad36e
	crc_rpc_answer_dropped_running = 0xcd78e586
	crc_rpc_answer_dropped         = 0xa43ad8b7
	crc_future_salt                = 0x0949d9dc
	crc_future_salts               = 0xae500895
	crc_pong                       = 0x347773c5
	crc_destroy_session_ok         = 0xe22045fc
	crc_destroy_session_none       = 0x62d350c9
	crc_new_session_created        = 0x9ec20908
	crc_msg_container              = 0x73f1f8dc
	crc_msg_copy                   = 0xe06046b2
	crc_gzip_packed                = 0x3072cfa1
	crc_msgs_ack                   = 0x62d6b459
	crc_bad_msg_notification       = 0xa7eff811
	crc_bad_server_salt            = 0xedab447b
	crc_msg_resend_req             = 0x7d861a08
	crc_msgs_state_req             = 0xda69fb52
	crc_msgs_state_info            = 0x04deb57d
	crc_msgs_all_info              = 0x8cc0d131
	crc_msg_detailed_info          = 0x276d3ec6
	crc_msg_new_detailed_info      = 0x809db6df
	crc_req_pq                     = 0x60469778
	crc_req_DH_params              = 0xd712e4be
	crc_set_client_DH_params       = 0xf5045f1f
	crc_rpc_drop_answer            = 0x58e4a740
	crc_get_future_salts           = 0xb921bd04
	crc_ping                       = 0x7abe77ec
	crc_ping_delay_disconnect      = 0xf3427b8c
	crc_destroy_session            = 0xe7512126
	crc_http_wait                  = 0x9299359f
)
