package sia

func (s *Sia) AddUInt8(n uint8) *Sia {
	s.Content = append(s.Content, byte(n))
	return s
}

func (s *Sia) ReadUInt8() uint8 {
	if s.Index+1 > uint64(len(s.Content)) {
		return 0
	}

	n := uint8(s.Content[s.Index])
	s.Index += 1
	return n
}
