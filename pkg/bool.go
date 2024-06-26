package sia

func (s *sia) AddBool(b bool) Sia {
	if b {
		s.Content = append(s.Content, 1)
	} else {
		s.Content = append(s.Content, 0)
	}
	return s
}

func (s *sia) ReadBool() bool {
	if s.Index >= uint64(len(s.Content)) {
		return false
	}

	if s.Content[s.Index] == 0 {
		s.Index++
		return false
	}

	s.Index++
	return true
}
