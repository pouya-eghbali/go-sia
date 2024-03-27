package sia

func (s *Sia) AddInt8(n int8) *Sia {
	s.Content = append(s.Content, byte(n))
	return s
}

func (s *Sia) ReadInt8() int8 {
	if s.Index+1 > uint64(len(s.Content)) {
		return 0
	}

	n := int8(s.Content[s.Index])
	s.Index += 1
	return n
}
