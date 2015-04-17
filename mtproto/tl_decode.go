package mtproto

import (
	"encoding/binary"
	"errors"
	"fmt"
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

func (m *DecodeBuf) Long() (r int64) {
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

func (m *DecodeBuf) Int() (r int32) {
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

func (m *DecodeBuf) UInt() (r uint32) {
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

func (m *DecodeBuf) Bytes(size int) (r []byte) {
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

func (m *DecodeBuf) StringBytes() (r []byte) {
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

func (m *DecodeBuf) String() (r string) {
	b := m.StringBytes()
	if m.err != nil {
		return ""
	}
	x := string(b)
	return x
}

func (m *DecodeBuf) BigInt() (r *big.Int) {
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

func (m *DecodeBuf) VectorLong() (r []int64) {
	constructor := m.UInt()
	if m.err != nil {
		return nil
	}
	if constructor != crc_vector {
		m.err = errors.New("DecodeVectorLong: Wrong constructor")
		return nil
	}
	size := m.Int()
	if m.err != nil {
		return nil
	}
	if size <= 0 {
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

func (m *DecodeBuf) Bool() (r bool) {
	constructor := m.UInt()
	if m.err != nil {
		return false
	}
	switch constructor {
	case crc_bool_false:
		return false
	case crc_bool_true:
		return true
	}
	return false
}

func (m *DecodeBuf) Vector(level int) []interface{} {
	constructor := m.UInt()
	if m.err != nil {
		return nil
	}
	if constructor != crc_vector {
		m.err = errors.New("DecodeVector: Wrong constructor")
		return nil
	}
	size := m.Int()
	if m.err != nil {
		return nil
	}
	if size <= 0 {
		m.err = errors.New("DecodeVector: Wrong size")
		return nil
	}
	x := make([]interface{}, size)
	i := int32(0)
	for i < size {
		y := m.Object(level)
		if m.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	return x
}

func (m *DecodeBuf) Object(level int) (r interface{}) {
	constructor := m.UInt()
	if m.err != nil {
		return nil
	}

	// fmt.Printf("[%08x]\n", constructor)

	switch constructor {

	case crc_resPQ:
		r = &TL_resPQ{m.Bytes(16), m.Bytes(16), m.BigInt(), m.VectorLong()}

	case crc_server_DH_params_ok:
		r = &TL_server_DH_params_ok{m.Bytes(16), m.Bytes(16), m.StringBytes()}

	case crc_server_DH_inner_data:
		r = &TL_server_DH_inner_data{
			m.Bytes(16), m.Bytes(16), m.Int(),
			m.BigInt(), m.BigInt(), m.Int(),
		}

	case crc_dh_gen_ok:
		r = &TL_dh_gen_ok{m.Bytes(16), m.Bytes(16), m.Bytes(16)}

	case crc_ping:
		r = &TL_ping{m.Long()}

	case crc_pong:
		r = &TL_pong{m.Long(), m.Long()}

	case crc_msg_container:
		size := m.Int()
		arr := make([]TL_message, size)
		for i := int32(0); i < size; i++ {
			arr[i] = TL_message{m.Long(), m.Int(), m.Int(), m.Object(level + 1)}
			if m.err != nil {
				return nil
			}
		}
		r = arr

	case crc_rpc_result:
		r = &TL_rpc_result{m.Long(), m.Object(level + 1)}

	case crc_new_session_created:
		r = &TL_new_session_created{m.Long(), m.Long(), m.Bytes(8)}

	case crc_bad_server_salt:
		r = &TL_bad_server_salt{m.Long(), m.Int(), m.Int(), m.Bytes(8)}

	case crc_msgs_ack:
		r = &TL_msgs_ack{m.VectorLong()}

	case crc_config:
		r = &TL_config{
			m.Int(),
			m.Bool(),
			m.Int(),
			func() []TL_dcOption {
				x := m.Vector(level + 1)
				y := make([]TL_dcOption, len(x))
				for i, v := range x {
					y[i] = *(v.(*TL_dcOption))
				}
				return y
			}(),
			m.Int(),
		}

	case crc_dcOption:
		r = &TL_dcOption{m.Int(), m.String(), m.String(), m.Int()}

	default:
		m.err = fmt.Errorf("Unknown constructor: %08x", constructor)
		return nil

	}

	if m.err != nil {
		return nil
	}

	return
}
