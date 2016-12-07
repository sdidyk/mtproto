package mtproto

import (
	"crypto/aes"
	"crypto/rsa"
	sha1lib "crypto/sha1"
	"errors"
	"math/big"
	"math/rand"
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

func sha1(data []byte) []byte {
	r := sha1lib.Sum(data)
	return r[:]
}

func RSA_encrypt(em []byte) []byte {
	z := make([]byte, 255)
	copy(z, em)

	c := new(big.Int)
	c.Exp(new(big.Int).SetBytes(z), big.NewInt(int64(telegramPublicKey.E)), telegramPublicKey.N)

	res := make([]byte, 256)
	copy(res, c.Bytes())

	return res
}

func SplitPQ(pq *big.Int) (p1, p2 *big.Int) {
	value_0 := big.NewInt(0)
	value_1 := big.NewInt(1)
	value_15 := big.NewInt(15)
	value_17 := big.NewInt(17)
	rndmax := big.NewInt(0).SetBit(big.NewInt(0), 64, 1)

	what := big.NewInt(0).Set(pq)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	g := big.NewInt(0)
	i := 0
	for !(g.Cmp(value_1) == 1 && g.Cmp(what) == -1) {
		q := big.NewInt(0).Rand(rnd, rndmax)
		q = q.And(q, value_15)
		q = q.Add(q, value_17)
		q = q.Mod(q, what)

		x := big.NewInt(0).Rand(rnd, rndmax)
		whatnext := big.NewInt(0).Sub(what, value_1)
		x = x.Mod(x, whatnext)
		x = x.Add(x, value_1)

		y := big.NewInt(0).Set(x)
		lim := 1 << (uint(i) + 18)
		j := 1
		flag := true

		for j < lim && flag {
			a := big.NewInt(0).Set(x)
			b := big.NewInt(0).Set(x)
			c := big.NewInt(0).Set(q)

			for b.Cmp(value_0) == 1 {
				b2 := big.NewInt(0)
				if b2.And(b, value_1).Cmp(value_0) == 1 {
					c.Add(c, a)
					if c.Cmp(what) >= 0 {
						c.Sub(c, what)
					}
				}
				a.Add(a, a)
				if a.Cmp(what) >= 0 {
					a.Sub(a, what)
				}
				b.Rsh(b, 1)
			}
			x.Set(c)

			z := big.NewInt(0)
			if x.Cmp(y) == -1 {
				z.Add(what, x)
				z.Sub(z, y)
			} else {
				z.Sub(x, y)
			}
			g.GCD(nil, nil, z, what)

			if (j & (j - 1)) == 0 {
				y.Set(x)
			}
			j = j + 1

			if g.Cmp(value_1) != 0 {
				flag = false
			}
		}
		i = i + 1
	}

	p1 = big.NewInt(0).Set(g)
	p2 = big.NewInt(0).Div(what, g)

	if p1.Cmp(p2) == 1 {
		p1, p2 = p2, p1
	}

	return
}

func MakeGAB(g int32, g_a, dh_prime *big.Int) (b, g_b, g_ab *big.Int) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	rndmax := big.NewInt(0).SetBit(big.NewInt(0), 2048, 1)
	b = big.NewInt(0).Rand(rnd, rndmax)
	g_b = big.NewInt(0).Exp(big.NewInt(int64(g)), b, dh_prime)
	g_ab = big.NewInt(0).Exp(g_a, b, dh_prime)

	return
}

func generateAES(msg_key, auth_key []byte, decode bool) ([]byte, []byte) {
	var x int
	if decode {
		x = 8
	} else {
		x = 0
	}
	aes_key := make([]byte, 0, 32)
	aes_iv := make([]byte, 0, 32)
	t_a := make([]byte, 0, 48)
	t_b := make([]byte, 0, 48)
	t_c := make([]byte, 0, 48)
	t_d := make([]byte, 0, 48)

	t_a = append(t_a, msg_key...)
	t_a = append(t_a, auth_key[x:x+32]...)

	t_b = append(t_b, auth_key[32+x:32+x+16]...)
	t_b = append(t_b, msg_key...)
	t_b = append(t_b, auth_key[48+x:48+x+16]...)

	t_c = append(t_c, auth_key[64+x:64+x+32]...)
	t_c = append(t_c, msg_key...)

	t_d = append(t_d, msg_key...)
	t_d = append(t_d, auth_key[96+x:96+x+32]...)

	sha1_a := sha1(t_a)
	sha1_b := sha1(t_b)
	sha1_c := sha1(t_c)
	sha1_d := sha1(t_d)

	aes_key = append(aes_key, sha1_a[0:8]...)
	aes_key = append(aes_key, sha1_b[8:8+12]...)
	aes_key = append(aes_key, sha1_c[4:4+12]...)

	aes_iv = append(aes_iv, sha1_a[8:8+12]...)
	aes_iv = append(aes_iv, sha1_b[0:8]...)
	aes_iv = append(aes_iv, sha1_c[16:16+4]...)
	aes_iv = append(aes_iv, sha1_d[0:8]...)

	return aes_key, aes_iv
}

func AES256IGE_encrypt(data, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(data) < aes.BlockSize {
		return nil, errors.New("AES256IGE: data too small to encrypt")
	}
	if len(data)%aes.BlockSize != 0 {
		return nil, errors.New("AES256IGE: data not divisible by block size")
	}

	t := make([]byte, aes.BlockSize)
	x := make([]byte, aes.BlockSize)
	y := make([]byte, aes.BlockSize)
	copy(x, iv[:aes.BlockSize])
	copy(y, iv[aes.BlockSize:])
	encrypted := make([]byte, len(data))

	i := 0
	for i < len(data) {
		xor(x, data[i:i+aes.BlockSize])
		block.Encrypt(t, x)
		xor(t, y)
		x, y = t, data[i:i+aes.BlockSize]
		copy(encrypted[i:], t)
		i += aes.BlockSize
	}

	return encrypted, nil
}

func AES256IGE_decrypt(data, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(data) < aes.BlockSize {
		return nil, errors.New("AES256IGE: data too small to decrypt")
	}
	if len(data)%aes.BlockSize != 0 {
		return nil, errors.New("AES256IGE: data not divisible by block size")
	}

	t := make([]byte, aes.BlockSize)
	x := make([]byte, aes.BlockSize)
	y := make([]byte, aes.BlockSize)
	copy(x, iv[:aes.BlockSize])
	copy(y, iv[aes.BlockSize:])
	decrypted := make([]byte, len(data))

	i := 0
	for i < len(data) {
		xor(y, data[i:i+aes.BlockSize])
		block.Decrypt(t, y)
		xor(t, x)
		y, x = t, data[i:i+aes.BlockSize]
		copy(decrypted[i:], t)
		i += aes.BlockSize
	}

	return decrypted, nil

}

func xor(dst, src []byte) {
	for i := range dst {
		dst[i] = dst[i] ^ src[i]
	}
}
