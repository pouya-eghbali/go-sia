package sia

func (s *Sia) AddString8(str string) *Sia {
	s.AddByteArray8([]byte(str))
	return s
}

func (s *Sia) ReadString8() string {
	return string(s.ReadByteArray8())
}

func (s *Sia) AddString16(str string) *Sia {
	s.AddByteArray16([]byte(str))
	return s
}

func (s *Sia) ReadString16() string {
	return string(s.ReadByteArray16())
}

func (s *Sia) AddString32(str string) *Sia {
	s.AddByteArray32([]byte(str))
	return s
}

func (s *Sia) ReadString32() string {
	return string(s.ReadByteArray32())
}

func (s *Sia) AddString64(str string) *Sia {
	s.AddByteArray64([]byte(str))
	return s
}

func (s *Sia) ReadString64() string {
	return string(s.ReadByteArray64())
}
