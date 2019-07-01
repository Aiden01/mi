package mi

import (
	"bytes"
	"encoding/binary"
)

func uvarint(x uint) []byte {
	if x == 0 {
		return []byte{0}
	}

	b := make([]byte, binary.MaxVarintLen64)
	i := len(b)
	for x > 0 {
		i--
		b[i] = byte(x&127) | 128
		x >>= 7
	}
	b[len(b)-1] &= 127
	return b[:i]
}

func binn(a interface{}) []byte {
	buf := bytes.NewBuffer(nil)
	binary.Write(buf, binary.BigEndian, a)
	return buf.Bytes()
}
