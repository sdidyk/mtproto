package mtproto

import (
	// "crypto/aes"
	// "crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/binary"
	"math/big"
	"time"
)

const (
	telegramPublicKey_N  = "24403446649145068056824081744112065346446136066297307473868293895086332508101251964919587745984311372853053253457835208829824428441874946556659953519213382748319518214765985662663680818277989736779506318868003755216402538945900388706898101286548187286716959100102939636333452457308619454821845196109544157601096359148241435922125602449263164512290854366930013825808102403072317738266383237191313714482187326643144603633877219028262697593882410403273959074350849923041765639673335775605842311578109726403165298875058941765362622936097839775380070572921007586266115476975819175319995527916042178582540628652481530373407"
	telegramPublicKey_E  = 65537
	telegramPublicKey_FP = 14101943622620965665
)

var telegramPublicKey rsa.PublicKey

func init() {
	telegramPublicKey.N, _ = new(big.Int).SetString(telegramPublicKey_N, 10)
	telegramPublicKey.E = telegramPublicKey_E
}

func Sha1(data []byte) []byte {
	r := sha1.Sum(data)
	return r[:]
}

func RSAEncode(em []byte) []byte {
	z := make([]byte, 255)
	copy(z, em)

	c := new(big.Int)
	c.Exp(new(big.Int).SetBytes(z), big.NewInt(int64(telegramPublicKey.E)), telegramPublicKey.N)

	res := make([]byte, 256)
	copy(res, c.Bytes())

	return res
}

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
	x = append(x, EncodeUInt(req_pq)...)
	x = append(x, EncodeBytes(nonce)...)
	return x
}

func Encode_TL_p_q_inner_data(pq, p, q *big.Int, nonce, server_nonce, new_nonce []byte) []byte {
	x := make([]byte, 0, 256)
	x = append(x, EncodeUInt(p_q_inner_data)...)
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
	x = append(x, EncodeUInt(req_DH_params)...)
	x = append(x, EncodeBytes(nonce)...)
	x = append(x, EncodeBytes(server_nonce)...)
	x = append(x, EncodeBigInt(p)...)
	x = append(x, EncodeBigInt(q)...)
	x = append(x, EncodeLong(int64(fp))...)
	x = append(x, EncodeStringBytes(encdata)...)
	return x
}
