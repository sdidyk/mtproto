package mtproto

import (
	"crypto/rand"
	"encoding/binary"
	"math/big"
	"time"
)

func GenerateNonce(size int) []byte {
	b := make([]byte, size)
	rand.Read(b)
	return b
}

func GenerateMessageId() int64 {
	const nano = 1000 * 1000 * 1000
	time := time.Now().UnixNano()

	return ((time / nano) << 32) | ((time % nano) & -4)
}

func EncodeInt(s int32) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(s))
	return bs
}

func EncodeUInt(s uint32) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, s)
	return bs
}

func EncodeLong(s int64) []byte {
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, uint64(s))
	return bs
}

func EncodeString(s string) []byte {
	return EncodeStringBytes([]byte(s))
}

func EncodeBigInt(s *big.Int) []byte {
	return EncodeStringBytes(s.Bytes())
}

func EncodeStringBytes(s []byte) []byte {
	var res []byte

	len := len(s)

	if len < 254 {
		nl := 1 + len + (4-(len+1)%4)&3
		res = make([]byte, nl)
		res[0] = byte(len)
		copy(res[1:], s)

	} else {
		nl := 4 + len + (4-len%4)&3
		res = make([]byte, nl)
		copy(res, EncodeInt(int32(len<<8|254)))
		copy(res[4:], s)

	}
	return res

}

func EncodeBytes(s []byte) []byte {
	return s
}

func Encode_TL_req_pq(nonce []byte) []byte {
	x := make([]byte, 0, 20)
	x = append(x, EncodeUInt(crc_req_pq)...)
	x = append(x, EncodeBytes(nonce)...)
	return x
}

func Encode_TL_p_q_inner_data(pq, p, q *big.Int, nonce, server_nonce, new_nonce []byte) []byte {
	x := make([]byte, 0, 256)
	x = append(x, EncodeUInt(crc_p_q_inner_data)...)
	x = append(x, EncodeBigInt(pq)...)
	x = append(x, EncodeBigInt(p)...)
	x = append(x, EncodeBigInt(q)...)
	x = append(x, EncodeBytes(nonce)...)
	x = append(x, EncodeBytes(server_nonce)...)
	x = append(x, EncodeBytes(new_nonce)...)
	return x
}

func Encode_TL_req_DH_params(nonce, server_nonce []byte, p, q *big.Int, fp uint64, encdata []byte) []byte {
	x := make([]byte, 0, 512)
	x = append(x, EncodeUInt(crc_req_DH_params)...)
	x = append(x, EncodeBytes(nonce)...)
	x = append(x, EncodeBytes(server_nonce)...)
	x = append(x, EncodeBigInt(p)...)
	x = append(x, EncodeBigInt(q)...)
	x = append(x, EncodeLong(int64(fp))...)
	x = append(x, EncodeStringBytes(encdata)...)
	return x
}

func Encode_TL_client_DH_inner_data(nonce, server_nonce []byte, retry int64, g_b *big.Int) []byte {
	x := make([]byte, 0, 256)
	x = append(x, EncodeUInt(crc_client_DH_inner_data)...)
	x = append(x, EncodeBytes(nonce)...)
	x = append(x, EncodeBytes(server_nonce)...)
	x = append(x, EncodeLong(retry)...)
	x = append(x, EncodeBigInt(g_b)...)
	return x
}

func Encode_TL_set_client_DH_params(nonce, server_nonce, encdata []byte) []byte {
	x := make([]byte, 0, 256)
	x = append(x, EncodeUInt(crc_set_client_DH_params)...)
	x = append(x, EncodeBytes(nonce)...)
	x = append(x, EncodeBytes(server_nonce)...)
	x = append(x, EncodeStringBytes(encdata)...)
	return x
}

func Encode_TL_ping(ping_id int64) []byte {
	x := make([]byte, 0, 32)
	x = append(x, EncodeUInt(crc_ping)...)
	x = append(x, EncodeLong(ping_id)...)
	return x
}

func Encode_TL_pong(msg_id, ping_id int64) []byte {
	x := make([]byte, 0, 32)
	x = append(x, EncodeUInt(crc_pong)...)
	x = append(x, EncodeLong(msg_id)...)
	x = append(x, EncodeLong(ping_id)...)
	return x
}

func Encode_TL_help_getConfig() []byte {
	x := make([]byte, 0, 8)
	x = append(x, EncodeUInt(crc_help_getConfig)...)
	return x
}
