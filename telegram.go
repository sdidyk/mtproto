package main

import (
	"encoding/hex"
	"github.com/sdidyk/telegram/mtproto"
	"math/big"
)

func main() {
	var x, y []byte

	x = append(x, mtproto.EncodeUInt(0x83c95aec)...)

	pq, _ := hex.DecodeString("17ED48941A08F981")
	x = append(x, mtproto.EncodeBigInt(new(big.Int).SetBytes(pq))...)

	p, _ := hex.DecodeString("494C553B")
	x = append(x, mtproto.EncodeBigInt(new(big.Int).SetBytes(p))...)

	q, _ := hex.DecodeString("53911073")
	x = append(x, mtproto.EncodeBigInt(new(big.Int).SetBytes(q))...)

	r4, _ := hex.DecodeString("3E0549828CCA27E966B301A48FECE2FCA5CF4D33F4A11EA877BA4AA573907330311C85DB234AA2640AFC4A76A735CF5B1F0FD68BD17FA181E1229AD867CC024D")
	x = append(x, mtproto.EncodeBytes(r4)...)

	y = append(y, mtproto.Sha1(x)...)
	y = append(y, x...)

	// res := mtproto.RSAEncode(y)
	// fmt.Println("[encoded]", hex.EncodeToString(res))

	m := new(mtproto.MTProto)
	m.Connect("149.154.175.50:443")
	m.Handshake()
	m.Dump()
}
