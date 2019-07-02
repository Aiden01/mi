package mi

import (
	"bytes"
	"encoding/binary"
)

type Track struct {
	Hits []*Hit
	BPM  uint
}

func (t *Track) MarshalBinary() []byte {
	buf := bytes.NewBuffer(nil)
	buf.Write(t.encodeHeaderChunk())
	buf.Write(t.encodeMetaChunk())
	buf.Write(t.encodeHits())
	return buf.Bytes()
}

func (*Track) encodeHeaderChunk() []byte {
	buf := bytes.NewBuffer(nil)
	buf.Write([]byte("MThd"))
	binary.Write(buf, binary.BigEndian, uint32(6))
	binary.Write(buf, binary.BigEndian, uint16(1))
	binary.Write(buf, binary.BigEndian, uint16(2))
	binary.Write(buf, binary.BigEndian, uint16(96))

	return buf.Bytes()
}

func (t *Track) encodeMetaChunk() []byte {
	mpb := 1 / float64(t.BPM)
	uspb := uint32(mpb * 60 * 1000000)

	buf := bytes.NewBuffer(nil)
	buf.Write([]byte("MTrk"))

	buf2 := bytes.NewBuffer(nil)

	buf2.Write([]byte{0, 0xFF, 0x58, 4, 4, 2, 24, 8})
	buf2.Write([]byte{0, 0xFF, 0x51, 3})
	buf2.Write(bin(uspb)[1:])
	buf2.Write([]byte{0, 0xFF, 0x2F, 0})

	buf.Write(bin(uint32(buf2.Len())))
	return append(buf.Bytes(), buf2.Bytes()...)
}

func (t *Track) encodeHits() []byte {
	buf := bytes.NewBuffer([]byte("MTrk"))
	buf2 := bytes.NewBuffer(nil)
	for _, h := range t.Hits {
		buf2.Write(h.encode())
	}
	buf2.Write([]byte{0, 0xFF, 0x2F, 0})
	buf.Write(bin(uint32(buf2.Len())))

	return append(buf.Bytes(), buf2.Bytes()...)
}

type Hit struct {
	Notes map[byte]Velocity   // Notes to strike with their velocities
	T     uint                // Nimber of ticks this hit lasts (96 is a quarter bar)
}

func (h *Hit) encode() []byte {
	buf := bytes.NewBuffer(nil)
	for n, v:= range h.Notes {
		buf.Write([]byte{0, 0x99, n, byte(v)})
	}
	first := true
	for n:= range h.Notes {
		if first {
			buf.Write(uvarint(h.T))
			first = false
		} else {
			buf.Write(uvarint(0))
		}
		buf.Write([]byte{0x89, n, 64})
	}
	return buf.Bytes()
}

type Velocity byte

const (
	PPP Velocity = 16  // Pianississimo
	PP  Velocity = 32  // Pianissimo
	P   Velocity = 48  // Piano
	MP  Velocity = 64  // Mezzo-piano
	MF  Velocity = 80  // Mezzo-forte
	F   Velocity = 96  // Forte
	FF  Velocity = 112 // Fortissimo
	FFF Velocity = 127 // Fortississimo
)
