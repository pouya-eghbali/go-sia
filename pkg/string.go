package sia

func (s *sia) AddString(str string) Sia {
	s.AddByteArray([]byte(str))
	return s
}

func (s *sia) ReadString(length uint64) string {
	return string(s.ReadByteArray(length))
}

func (s *sia) AddString8(str string) Sia {
	s.AddByteArray8([]byte(str))
	return s
}

func (s *sia) ReadString8() string {
	return string(s.ReadByteArray8())
}

func (s *sia) AddString16(str string) Sia {
	s.AddByteArray16([]byte(str))
	return s
}

func (s *sia) ReadString16() string {
	return string(s.ReadByteArray16())
}

func (s *sia) AddString32(str string) Sia {
	s.AddByteArray32([]byte(str))
	return s
}

func (s *sia) ReadString32() string {
	return string(s.ReadByteArray32())
}

func (s *sia) AddString64(str string) Sia {
	s.AddByteArray64([]byte(str))
	return s
}

func (s *sia) ReadString64() string {
	return string(s.ReadByteArray64())
}
