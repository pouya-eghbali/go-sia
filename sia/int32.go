package sia

import "encoding/binary"

func (s *Sia) AddInt32(n int32) *Sia {
	s.Content = binary.LittleEndian.AppendUint32(s.Content, uint32(n))
	return s
}

func (s *Sia) ReadInt32() int32 {
	if s.Index+4 > uint64(len(s.Content)) {
		return 0
	}

	n := int32(binary.LittleEndian.Uint32(s.Content[s.Index : s.Index+4]))
	s.Index += 4
	return n
}
