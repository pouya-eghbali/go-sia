package sia

import "encoding/binary"

func (s *Sia) AddInt16(n int16) *Sia {
	s.Content = binary.LittleEndian.AppendUint16(s.Content, uint16(n))
	return s
}

func (s *Sia) ReadInt16() int16 {
	if s.Index+2 > uint64(len(s.Content)) {
		return 0
	}

	n := int16(binary.LittleEndian.Uint16(s.Content[s.Index : s.Index+2]))
	s.Index += 2
	return n
}
