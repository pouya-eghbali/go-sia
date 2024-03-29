package sia

import "encoding/binary"

func (s *sia) AddUInt8(n uint8) Sia {
	s.Content = append(s.Content, byte(n))
	return s
}

func (s *sia) ReadUInt8() uint8 {
	if s.Index+1 > uint64(len(s.Content)) {
		return 0
	}

	n := uint8(s.Content[s.Index])
	s.Index += 1
	return n
}

func (s *sia) AddUInt64(n uint64) Sia {
	s.Content = binary.LittleEndian.AppendUint64(s.Content, n)
	return s
}

func (s *sia) ReadUInt64() uint64 {
	if s.Index+8 > uint64(len(s.Content)) {
		return 0
	}

	n := binary.LittleEndian.Uint64(s.Content[s.Index : s.Index+8])
	s.Index += 8
	return n
}

func (s *sia) AddUInt32(n uint32) Sia {
	s.Content = binary.LittleEndian.AppendUint32(s.Content, n)
	return s
}

func (s *sia) ReadUInt32() uint32 {
	if s.Index+4 > uint64(len(s.Content)) {
		return 0
	}

	n := binary.LittleEndian.Uint32(s.Content[s.Index : s.Index+4])
	s.Index += 2
	return n
}

func (s *sia) AddUInt16(n uint16) Sia {
	s.Content = binary.LittleEndian.AppendUint16(s.Content, n)
	return s
}

func (s *sia) ReadUInt16() uint16 {
	if s.Index+2 > uint64(len(s.Content)) {
		return 0
	}

	n := binary.LittleEndian.Uint16(s.Content[s.Index : s.Index+2])
	s.Index += 2
	return n
}
