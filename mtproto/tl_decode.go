package mtproto

import (
	"encoding/binary"
	"errors"
	"math/big"
)

func (m *MTProto) DecodeLong() (r int64, err error) {
	if m.off+8 > m.size {
		return 0, errors.New("DecodeLong: короткий пакет")
	}
	x := int64(binary.LittleEndian.Uint64(m.buf[m.off : m.off+8]))
	m.off += 8
	return x, nil
}

func (m *MTProto) DecodeInt() (r int32, err error) {
	if m.off+4 > m.size {
		return 0, errors.New("DecodeInt: короткий пакет")
	}
	x := binary.LittleEndian.Uint32(m.buf[m.off : m.off+4])
	m.off += 4
	return int32(x), nil
}

func (m *MTProto) DecodeUInt() (r uint32, err error) {
	if m.off+4 > m.size {
		return 0, errors.New("DecodeUInt: короткий пакет")
	}
	x := binary.LittleEndian.Uint32(m.buf[m.off : m.off+4])
	m.off += 4
	return x, nil
}

func (m *MTProto) DecodeBytes(size int) (r []byte, err error) {
	if m.off+size > m.size {
		return nil, errors.New("DecodeBytes: короткий пакет")
	}
	x := make([]byte, size)
	copy(x, m.buf[m.off:m.off+size])
	m.off += size
	return x, nil
}

func (m *MTProto) DecodeStringBytes() (r []byte, err error) {
	var size, padding int

	if m.off+1 > m.size {
		return nil, errors.New("DecodeStringBytes: короткий пакет")
	}
	size = int(m.buf[m.off])
	m.off++
	padding = (4 - ((size + 1) % 4)) & 3
	if size == 254 {
		if m.off+3 > m.size {
			return nil, errors.New("DecodeStringBytes: короткий пакет")
		}
		size = int(m.buf[m.off]) | int(m.buf[m.off+1])<<8 | int(m.buf[m.off+2])<<16
		m.off += 3
		padding = (4 - size%4) & 3
	}

	if m.off+size > m.size {
		return nil, errors.New("DecodeStringBytes: короткий пакет (size)")
	}
	x := make([]byte, size)
	copy(x, m.buf[m.off:m.off+size])
	m.off += size

	if m.off+padding > m.size {
		return nil, errors.New("DecodeStringBytes: короткий пакет (padding)")
	}
	m.off += padding

	return x, nil
}

func (m *MTProto) DecodeBigInt() (r *big.Int, err error) {
	b, err := m.DecodeStringBytes()
	if err != nil {
		return nil, err
	}
	y := make([]byte, len(b)+1)
	y[0] = 0
	copy(y[1:], b)
	x := new(big.Int).SetBytes(y)
	return x, nil
}

func (m *MTProto) DecodeVectorLong() (r []int64, err error) {
	constructor, err := m.DecodeUInt()
	if err != nil {
		return nil, err
	}
	if constructor != crc_vector {
		return nil, errors.New("DecodeVectorLong: Неправильный конструктор")
	}
	size, err := m.DecodeInt()
	if err != nil {
		return nil, err
	}
	if size <= 0 {
		return nil, errors.New("DecodeVectorLong: Неправильный размер")
	}
	x := make([]int64, size)
	i := int32(0)
	for i < size {
		y, err := m.DecodeLong()
		if err != nil {
			return nil, err
		}
		x[i] = y
		i++
	}
	return x, nil
}
