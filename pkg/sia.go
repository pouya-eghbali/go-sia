package sia

import (
	"math/big"
)

type Sia interface {
	Seek(index uint64) Sia
	Bytes() []byte
	Offset() uint64

	EmbedSia(s2 *sia) Sia
	EmbedBytes(b []byte) Sia

	AddBigInt(n *big.Int) Sia
	ReadBigInt() *big.Int
	AddBool(b bool) Sia
	ReadBool() bool
	AddUInt64(n uint64) Sia
	ReadUInt64() uint64
	AddUInt32(n uint32) Sia
	ReadUInt32() uint32
	AddUInt16(n uint16) Sia
	ReadUInt16() uint16
	AddUInt8(n uint8) Sia
	ReadUInt8() uint8
	AddStringN(str string) Sia
	ReadStringN(length uint64) string
	AddString8(str string) Sia
	ReadString8() string
	AddString16(str string) Sia
	ReadString16() string
	AddString32(str string) Sia
	ReadString32() string
	AddString64(str string) Sia
	ReadString64() string
	AddInt64(n int64) Sia
	ReadInt64() int64
	AddInt32(n int32) Sia
	ReadInt32() int32
	AddInt16(n int16) Sia
	ReadInt16() int16
	AddInt8(n int8) Sia
	ReadInt8() int8
	readBytesAt(start uint64, length uint64) []byte
	AddByteArrayN(bytes []byte) Sia
	ReadByteArrayN(length uint64) []byte
	AddByteArray8(bytes []byte) Sia
	ReadByteArray8() []byte
	AddByteArray16(bytes []byte) Sia
	ReadByteArray16() []byte
	AddByteArray32(bytes []byte) Sia
	ReadByteArray32() []byte
	AddByteArray64(bytes []byte) Sia
	ReadByteArray64() []byte
}

type sia struct {
	Content []byte
	Index   uint64
}

type ArraySia[T any] struct {
	sia
}

func (s *sia) Seek(index uint64) Sia {
	s.Index = index
	return s
}

func (s *sia) EmbedSia(embedSia *sia) Sia {
	s.Content = append(s.Content, embedSia.Content...)
	return s
}

func (s *sia) EmbedBytes(b []byte) Sia {
	s.Content = append(s.Content, b...)
	return s
}

func (s *sia) Bytes() []byte {
	return s.Content
}

func (s *sia) Offset() uint64 {
	return s.Index
}

func New() Sia {
	return &sia{}
}

func NewFromBytes(content []byte) Sia {
	return &sia{
		Content: content,
	}
}
