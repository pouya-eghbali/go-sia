package sia

func (s *Sia) readBytesAt(start uint64, length uint64) []byte {
	if start+length > uint64(len(s.Content)) {
		return []byte{}
	}

	bytes := s.Content[start : start+length]
	s.Index += length
	return bytes
}

func (s *Sia) AddByteArray8(bytes []byte) *Sia {
	s.AddUInt8(uint8(len(bytes)))
	s.Content = append(s.Content, bytes...)
	return s
}

func (s *Sia) ReadByteArray8() []byte {
	length := uint64(s.ReadUInt8())
	return s.readBytesAt(s.Index, length)
}

func (s *Sia) AddByteArray16(bytes []byte) *Sia {
	s.AddUInt16(uint16(len(bytes)))
	s.Content = append(s.Content, bytes...)
	return s
}

func (s *Sia) ReadByteArray16() []byte {
	length := uint64(s.ReadUInt16())
	return s.readBytesAt(s.Index, length)
}

func (s *Sia) AddByteArray32(bytes []byte) *Sia {
	s.AddUInt32(uint32(len(bytes)))
	s.Content = append(s.Content, bytes...)
	return s
}

func (s *Sia) ReadByteArray32() []byte {
	length := uint64(s.ReadUInt32())
	return s.readBytesAt(s.Index, length)
}

func (s *Sia) AddByteArray64(bytes []byte) *Sia {
	s.AddUInt64(uint64(len(bytes)))
	s.Content = append(s.Content, bytes...)
	return s
}

func (s *Sia) ReadByteArray64() []byte {
	length := s.ReadUInt64()
	return s.readBytesAt(s.Index, length)
}
