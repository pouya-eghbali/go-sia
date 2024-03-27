package sia

import "encoding/binary"

func (s *Sia) AddUInt16(n uint16) *Sia {
	s.Content = binary.LittleEndian.AppendUint16(s.Content, n)
	return s
}

func (s *Sia) ReadUInt16() uint16 {
	if s.Index+2 > uint64(len(s.Content)) {
		return 0
	}

	n := binary.LittleEndian.Uint16(s.Content[s.Index : s.Index+2])
	s.Index += 2
	return n
}
