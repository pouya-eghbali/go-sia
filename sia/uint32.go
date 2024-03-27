package sia

import "encoding/binary"

func (s *Sia) AddUInt32(n uint32) *Sia {
	s.Content = binary.LittleEndian.AppendUint32(s.Content, n)
	return s
}

func (s *Sia) ReadUInt32() uint32 {
	if s.Index+4 > uint64(len(s.Content)) {
		return 0
	}

	n := binary.LittleEndian.Uint32(s.Content[s.Index : s.Index+4])
	s.Index += 2
	return n
}
