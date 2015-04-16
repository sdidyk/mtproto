package mtproto

import (
	"encoding/binary"
	"errors"
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

func (m *DecodeBuf) DecodeLong() (r int64) {
	if m.err != nil {
		return 0
	}
	if m.off+8 > m.size {
		m.err = errors.New("DecodeLong: короткий пакет")
		return 0
	}
	x := int64(binary.LittleEndian.Uint64(m.buf[m.off : m.off+8]))
	m.off += 8
	return x
}

func (m *DecodeBuf) DecodeInt() (r int32) {
	if m.err != nil {
		return 0
	}
	if m.off+4 > m.size {
		m.err = errors.New("DecodeInt: короткий пакет")
		return 0
	}
	x := binary.LittleEndian.Uint32(m.buf[m.off : m.off+4])
	m.off += 4
	return int32(x)
}

func (m *DecodeBuf) DecodeUInt() (r uint32) {
	if m.err != nil {
		return 0
	}
	if m.off+4 > m.size {
		m.err = errors.New("DecodeUInt: короткий пакет")
		return 0
	}
	x := binary.LittleEndian.Uint32(m.buf[m.off : m.off+4])
	m.off += 4
	return x
}

func (m *DecodeBuf) DecodeBytes(size int) (r []byte) {
	if m.err != nil {
		return nil
	}
	if m.off+size > m.size {
		m.err = errors.New("DecodeBytes: короткий пакет")
		return nil
	}
	x := make([]byte, size)
	copy(x, m.buf[m.off:m.off+size])
	m.off += size
	return x
}

func (m *DecodeBuf) DecodeStringBytes() (r []byte) {
	if m.err != nil {
		return nil
	}
	var size, padding int

	if m.off+1 > m.size {
		m.err = errors.New("DecodeStringBytes: короткий пакет")
		return nil
	}
	size = int(m.buf[m.off])
	m.off++
	padding = (4 - ((size + 1) % 4)) & 3
	if size == 254 {
		if m.off+3 > m.size {
			m.err = errors.New("DecodeStringBytes: короткий пакет")
			return nil
		}
		size = int(m.buf[m.off]) | int(m.buf[m.off+1])<<8 | int(m.buf[m.off+2])<<16
		m.off += 3
		padding = (4 - size%4) & 3
	}

	if m.off+size > m.size {
		m.err = errors.New("DecodeStringBytes: короткий пакет (size)")
		return nil
	}
	x := make([]byte, size)
	copy(x, m.buf[m.off:m.off+size])
	m.off += size

	if m.off+padding > m.size {
		m.err = errors.New("DecodeStringBytes: короткий пакет (padding)")
		return nil
	}
	m.off += padding

	return x
}

func (m *DecodeBuf) DecodeBigInt() (r *big.Int) {
	b := m.DecodeStringBytes()
	if m.err != nil {
		return nil
	}
	y := make([]byte, len(b)+1)
	y[0] = 0
	copy(y[1:], b)
	x := new(big.Int).SetBytes(y)
	return x
}

func (m *DecodeBuf) DecodeVectorLong() (r []int64) {
	constructor := m.DecodeUInt()
	if m.err != nil {
		return nil
	}
	if constructor != crc_vector {
		m.err = errors.New("DecodeVectorLong: Неправильный конструктор")
		return nil
	}
	size := m.DecodeInt()
	if m.err != nil {
		return nil
	}
	if size <= 0 {
		m.err = errors.New("DecodeVectorLong: Неправильный размер")
		return nil
	}
	x := make([]int64, size)
	i := int32(0)
	for i < size {
		y := m.DecodeLong()
		if m.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	return x
}
