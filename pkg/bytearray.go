package sia

func (s *sia) readBytesAt(start uint64, length uint64) []byte {
	if start+length > uint64(len(s.Content)) {
		return []byte{}
	}

	bytes := s.Content[start : start+length]
	s.Index += length
	return bytes
}

func (s *sia) AddByteArray8(bytes []byte) Sia {
	s.AddUInt8(uint8(len(bytes)))
	s.Content = append(s.Content, bytes...)
	return s
}

func (s *sia) ReadByteArray8() []byte {
	length := uint64(s.ReadUInt8())
	return s.readBytesAt(s.Index, length)
}

func (s *sia) AddByteArray16(bytes []byte) Sia {
	s.AddUInt16(uint16(len(bytes)))
	s.Content = append(s.Content, bytes...)
	return s
}

func (s *sia) ReadByteArray16() []byte {
	length := uint64(s.ReadUInt16())
	return s.readBytesAt(s.Index, length)
}

func (s *sia) AddByteArray32(bytes []byte) Sia {
	s.AddUInt32(uint32(len(bytes)))
	s.Content = append(s.Content, bytes...)
	return s
}

func (s *sia) ReadByteArray32() []byte {
	length := uint64(s.ReadUInt32())
	return s.readBytesAt(s.Index, length)
}

func (s *sia) AddByteArray64(bytes []byte) Sia {
	s.AddUInt64(uint64(len(bytes)))
	s.Content = append(s.Content, bytes...)
	return s
}

func (s *sia) ReadByteArray64() []byte {
	length := s.ReadUInt64()
	return s.readBytesAt(s.Index, length)
}
