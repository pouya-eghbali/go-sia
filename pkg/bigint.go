package sia

import (
	"math/big"
)

func (s *Sia) AddBigInt(n *big.Int) *Sia {
	bytes := n.Bytes()
	s.Content = append(s.Content, byte(len(bytes)))
	s.Content = append(s.Content, n.Bytes()...)
	return s
}

func (s *Sia) ReadBigInt() *big.Int {
	if s.Index >= uint64(len(s.Content)) {
		return new(big.Int)
	}

	length := uint64(s.Content[s.Index])
	s.Index++

	if s.Index+length > uint64(len(s.Content)) {
		return new(big.Int)
	}

	bytes := s.Content[s.Index : s.Index+length]
	s.Index += uint64(length)
	return new(big.Int).SetBytes(bytes)
}
