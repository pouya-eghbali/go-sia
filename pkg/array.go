package sia

func (s *ArraySia[T]) loop(array []T, fn func(s *ArraySia[T], item T)) Sia {
	for _, item := range array {
		fn(s, item)
	}
	return &s.sia
}

func (s *ArraySia[T]) read(length uint64, fn func(s *ArraySia[T]) T) []T {
	array := make([]T, length)
	for i := uint64(0); i < length; i++ {
		array[i] = fn(s)
	}
	return array
}

func (s *ArraySia[T]) AddArray8(array []T, fn func(s *ArraySia[T], item T)) Sia {
	length := uint8(len(array))
	s.AddUInt8(length)
	return s.loop(array, fn)
}

func (s *ArraySia[T]) ReadArray8(fn func(s *ArraySia[T]) T) []T {
	length := uint8(s.ReadUInt8())
	return s.read(uint64(length), fn)
}

func (s *ArraySia[T]) AddArray16(array []T, fn func(s *ArraySia[T], item T)) Sia {
	length := uint16(len(array))
	s.AddUInt16(length)
	return s.loop(array, fn)
}

func (s *ArraySia[T]) ReadArray16(fn func(s *ArraySia[T]) T) []T {
	length := uint16(s.ReadUInt16())
	return s.read(uint64(length), fn)
}

func (s *ArraySia[T]) AddArray32(array []T, fn func(s *ArraySia[T], item T)) Sia {
	length := uint32(len(array))
	s.AddUInt32(length)
	return s.loop(array, fn)
}

func (s *ArraySia[T]) ReadArray32(fn func(s *ArraySia[T]) T) []T {
	length := s.ReadUInt32()
	return s.read(uint64(length), fn)
}

func (s *ArraySia[T]) AddArray64(array []T, fn func(s *ArraySia[T], item T)) Sia {
	length := uint64(len(array))
	s.AddUInt64(length)
	return s.loop(array, fn)
}

func (s *ArraySia[T]) ReadArray64(fn func(s *ArraySia[T]) T) []T {
	length := s.ReadUInt64()
	return s.read(length, fn)
}

func NewSiaArray[T any]() *ArraySia[T] {
	return &ArraySia[T]{}
}

func NewArrayFromBytes[T any](content []byte) *ArraySia[T] {
	return &ArraySia[T]{
		sia{Content: content},
	}
}
