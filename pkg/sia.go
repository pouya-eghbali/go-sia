package sia

type Sia struct {
	Content []byte
	Index   uint64
}

type ArraySia[T any] struct {
	Sia
}

func (s *Sia) Seek(index uint64) *Sia {
	s.Index = index
	return s
}

func (s *Sia) SetContent(content []byte) *Sia {
	s.Content = content
	return s
}

func (s *Sia) EmbedSia(s2 *Sia) *Sia {
	s.Content = append(s.Content, s2.Content...)
	return s
}

func (s *Sia) EmbedBytes(b []byte) *Sia {
	s.Content = append(s.Content, b...)
	return s
}
