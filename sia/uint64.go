package sia

import "encoding/binary"

func (s *Sia) AddUInt64(n uint64) *Sia {
	s.Content = binary.LittleEndian.AppendUint64(s.Content, n)
	return s
}

func (s *Sia) ReadUInt64() uint64 {
	if s.Index+8 > uint64(len(s.Content)) {
		return 0
	}

	n := binary.LittleEndian.Uint64(s.Content[s.Index : s.Index+8])
	s.Index += 8
	return n
}
