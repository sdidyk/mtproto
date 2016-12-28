package mtproto

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"math/big"
)

type DecodeBuf struct {
	buf  []byte
	off  int
	size int
	err  error
}

func NewDecodeBuf(b []byte) *DecodeBuf {
	return &DecodeBuf{b, 0, len(b), nil}
}

func (m *DecodeBuf) Long() int64 {
	if m.err != nil {
		return 0
	}
	if m.off+8 > m.size {
		m.err = errors.New("DecodeLong")
		return 0
	}
	x := int64(binary.LittleEndian.Uint64(m.buf[m.off : m.off+8]))
	m.off += 8
	return x
}

func (m *DecodeBuf) Double() float64 {
	if m.err != nil {
		return 0
	}
	if m.off+8 > m.size {
		m.err = errors.New("DecodeDouble")
		return 0
	}
	x := math.Float64frombits(binary.LittleEndian.Uint64(m.buf[m.off : m.off+8]))
	m.off += 8
	return x
}

func (m *DecodeBuf) Int() int32 {
	if m.err != nil {
		return 0
	}
	if m.off+4 > m.size {
		m.err = errors.New("DecodeInt")
		return 0
	}
	x := binary.LittleEndian.Uint32(m.buf[m.off : m.off+4])
	m.off += 4
	return int32(x)
}

func (m *DecodeBuf) UInt() uint32 {
	if m.err != nil {
		return 0
	}
	if m.off+4 > m.size {
		m.err = errors.New("DecodeUInt")
		return 0
	}
	x := binary.LittleEndian.Uint32(m.buf[m.off : m.off+4])
	m.off += 4
	return x
}

func (m *DecodeBuf) Bytes(size int) []byte {
	if m.err != nil {
		return nil
	}
	if m.off+size > m.size {
		m.err = errors.New("DecodeBytes")
		return nil
	}
	x := make([]byte, size)
	copy(x, m.buf[m.off:m.off+size])
	m.off += size
	return x
}

func (m *DecodeBuf) StringBytes() []byte {
	if m.err != nil {
		return nil
	}
	var size, padding int

	if m.off+1 > m.size {
		m.err = errors.New("DecodeStringBytes")
		return nil
	}
	size = int(m.buf[m.off])
	m.off++
	padding = (4 - ((size + 1) % 4)) & 3
	if size == 254 {
		if m.off+3 > m.size {
			m.err = errors.New("DecodeStringBytes")
			return nil
		}
		size = int(m.buf[m.off]) | int(m.buf[m.off+1])<<8 | int(m.buf[m.off+2])<<16
		m.off += 3
		padding = (4 - size%4) & 3
	}

	if m.off+size > m.size {
		m.err = errors.New("DecodeStringBytes: Wrong size")
		return nil
	}
	x := make([]byte, size)
	copy(x, m.buf[m.off:m.off+size])
	m.off += size

	if m.off+padding > m.size {
		m.err = errors.New("DecodeStringBytes: Wrong padding")
		return nil
	}
	m.off += padding

	return x
}

func (m *DecodeBuf) String() string {
	b := m.StringBytes()
	if m.err != nil {
		return ""
	}
	x := string(b)
	return x
}

func (m *DecodeBuf) BigInt() *big.Int {
	b := m.StringBytes()
	if m.err != nil {
		return nil
	}
	y := make([]byte, len(b)+1)
	y[0] = 0
	copy(y[1:], b)
	x := new(big.Int).SetBytes(y)
	return x
}

func (m *DecodeBuf) VectorInt() []int32 {
	constructor := m.UInt()
	if m.err != nil {
		return nil
	}
	if constructor != crc_vector {
		m.err = fmt.Errorf("DecodeVectorInt: Wrong constructor (0x%08x)", constructor)
		return nil
	}
	size := m.Int()
	if m.err != nil {
		return nil
	}
	if size < 0 {
		m.err = errors.New("DecodeVectorInt: Wrong size")
		return nil
	}
	x := make([]int32, size)
	i := int32(0)
	for i < size {
		y := m.Int()
		if m.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	return x
}

func (m *DecodeBuf) VectorLong() []int64 {
	constructor := m.UInt()
	if m.err != nil {
		return nil
	}
	if constructor != crc_vector {
		m.err = fmt.Errorf("DecodeVectorLong: Wrong constructor (0x%08x)", constructor)
		return nil
	}
	size := m.Int()
	if m.err != nil {
		return nil
	}
	if size < 0 {
		m.err = errors.New("DecodeVectorLong: Wrong size")
		return nil
	}
	x := make([]int64, size)
	i := int32(0)
	for i < size {
		y := m.Long()
		if m.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	return x
}

func (m *DecodeBuf) VectorString() []string {
	constructor := m.UInt()
	if m.err != nil {
		return nil
	}
	if constructor != crc_vector {
		m.err = fmt.Errorf("DecodeVectorString: Wrong constructor (0x%08x)", constructor)
		return nil
	}
	size := m.Int()
	if m.err != nil {
		return nil
	}
	if size < 0 {
		m.err = errors.New("DecodeVectorString: Wrong size")
		return nil
	}
	x := make([]string, size)
	i := int32(0)
	for i < size {
		y := m.String()
		if m.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	return x
}

func (m *DecodeBuf) Bool() bool {
	constructor := m.UInt()
	if m.err != nil {
		return false
	}
	switch constructor {
	case crc_boolFalse:
		return false
	case crc_boolTrue:
		return true
	}
	return false
}

func (m *DecodeBuf) Vector() []TL {
	constructor := m.UInt()
	if m.err != nil {
		return nil
	}
	if constructor != crc_vector {
		m.err = fmt.Errorf("DecodeVector: Wrong constructor (0x%08x)", constructor)
		return nil
	}
	size := m.Int()
	if m.err != nil {
		return nil
	}
	if size < 0 {
		m.err = errors.New("DecodeVector: Wrong size")
		return nil
	}
	x := make([]TL, size)
	i := int32(0)
	for i < size {
		y := m.Object()
		if m.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	return x
}

func (m *DecodeBuf) Object() (r TL) {
	constructor := m.UInt()
	if m.err != nil {
		return nil
	}

	// fmt.Printf("[%08x]\n", constructor)
	// m.dump()

	switch constructor {

	case crc_resPQ:
		r = TL_resPQ{m.Bytes(16), m.Bytes(16), m.BigInt(), m.VectorLong()}

	case crc_server_DH_params_ok:
		r = TL_server_DH_params_ok{m.Bytes(16), m.Bytes(16), m.StringBytes()}

	case crc_server_DH_inner_data:
		r = TL_server_DH_inner_data{
			m.Bytes(16), m.Bytes(16), m.Int(),
			m.BigInt(), m.BigInt(), m.Int(),
		}

	case crc_dh_gen_ok:
		r = TL_dh_gen_ok{m.Bytes(16), m.Bytes(16), m.Bytes(16)}

	case crc_ping:
		r = TL_ping{m.Long()}

	case crc_pong:
		r = TL_pong{m.Long(), m.Long()}

	case crc_msg_container:
		size := m.Int()
		arr := make([]TL_MT_message, size)
		for i := int32(0); i < size; i++ {
			arr[i] = TL_MT_message{m.Long(), m.Int(), m.Int(), m.Object()}
			if m.err != nil {
				return nil
			}
		}
		r = TL_msg_container{arr}

	case crc_rpc_result:
		r = TL_rpc_result{m.Long(), m.Object()}

	case crc_rpc_error:
		r = TL_rpc_error{m.Int(), m.String()}

	case crc_new_session_created:
		r = TL_new_session_created{m.Long(), m.Long(), m.Bytes(8)}

	case crc_bad_server_salt:
		r = TL_bad_server_salt{m.Long(), m.Int(), m.Int(), m.Bytes(8)}

	case crc_bad_msg_notification:
		r = TL_crc_bad_msg_notification{m.Long(), m.Int(), m.Int()}

	case crc_msgs_ack:
		r = TL_msgs_ack{m.VectorLong()}

	case crc_gzip_packed:
		obj := make([]byte, 0, 4096)

		var buf bytes.Buffer
		_, _ = buf.Write(m.StringBytes())
		gz, _ := gzip.NewReader(&buf)

		b := make([]byte, 4096)
		for true {
			n, _ := gz.Read(b)
			obj = append(obj, b[0:n]...)
			if n <= 0 {
				break
			}
		}
		d := NewDecodeBuf(obj)
		r = d.Object()

	default:
		r = m.ObjectGenerated(constructor)

	}

	if m.err != nil {
		return nil
	}

	return
}

func (d *DecodeBuf) dump() {
	fmt.Println(hex.Dump(d.buf[d.off:d.size]))
}

func toBool(x TL) bool {
	_, ok := x.(TL_boolTrue)
	return ok
}
