package sia

import "encoding/binary"

func (s *sia) AddInt8(n int8) Sia {
	s.Content = append(s.Content, byte(n))
	return s
}

func (s *sia) ReadInt8() int8 {
	if s.Index+1 > uint64(len(s.Content)) {
		return 0
	}

	n := int8(s.Content[s.Index])
	s.Index += 1
	return n
}

func (s *sia) AddInt16(n int16) Sia {
	s.Content = binary.LittleEndian.AppendUint16(s.Content, uint16(n))
	return s
}

func (s *sia) ReadInt16() int16 {
	if s.Index+2 > uint64(len(s.Content)) {
		return 0
	}

	n := int16(binary.LittleEndian.Uint16(s.Content[s.Index : s.Index+2]))
	s.Index += 2
	return n
}

func (s *sia) AddInt32(n int32) Sia {
	s.Content = binary.LittleEndian.AppendUint32(s.Content, uint32(n))
	return s
}

func (s *sia) ReadInt32() int32 {
	if s.Index+4 > uint64(len(s.Content)) {
		return 0
	}

	n := int32(binary.LittleEndian.Uint32(s.Content[s.Index : s.Index+4]))
	s.Index += 4
	return n
}

func (s *sia) AddInt64(n int64) Sia {
	s.Content = binary.LittleEndian.AppendUint64(s.Content, uint64(n))
	return s
}

func (s *sia) ReadInt64() int64 {
	if s.Index+8 > uint64(len(s.Content)) {
		return 0
	}

	n := int64(binary.LittleEndian.Uint64(s.Content[s.Index : s.Index+8]))
	s.Index += 8
	return n
}
