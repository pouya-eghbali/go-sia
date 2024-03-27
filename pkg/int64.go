package sia

import "encoding/binary"

func (s *Sia) AddInt64(n int64) *Sia {
	s.Content = binary.LittleEndian.AppendUint64(s.Content, uint64(n))
	return s
}

func (s *Sia) ReadInt64() int64 {
	if s.Index+8 > uint64(len(s.Content)) {
		return 0
	}

	n := int64(binary.LittleEndian.Uint64(s.Content[s.Index : s.Index+8]))
	s.Index += 8
	return n
}
