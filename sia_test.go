package go_sia__test

import (
	"github.com/google/go-cmp/cmp"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSerialize(t *testing.T) {
	var sampleUint16 uint16 = 182
	sampleString := "hello world"

	rawByte := sia.New().
		AddUInt16(sampleUint16).
		AddString64(sampleString).
		Bytes()

	deserialized := sia.NewFromBytes(rawByte)
	gotSampleUint16 := deserialized.ReadUInt16()
	assert.Equal(t, sampleUint16, gotSampleUint16, "should be equal")
	gotSampleString := deserialized.ReadString64()
	assert.Equal(t, sampleString, gotSampleString, "should be equal")
}

type person struct {
	Name string
	Age  uint8
}

func (p *person) Sia() sia.Sia {
	return sia.New().
		AddString8(p.Name).
		AddUInt8(p.Age)
}

func TestArraySerialize(t *testing.T) {

	sampleArray := []person{
		{Name: "joe", Age: 12},
		{Name: "jane", Age: 13},
		{Name: "john", Age: 14},
	}

	rawByte := sia.NewSiaArray[person]().
		AddArray8(sampleArray, func(s *sia.ArraySia[person], item person) {
			s.EmbedBytes(item.Sia().Bytes())
		}).Bytes()

	deserialized := sia.NewArrayFromBytes[person](rawByte).ReadArray8(func(s *sia.ArraySia[person]) person {
		return person{
			Name: s.ReadString8(),
			Age:  s.ReadUInt8(),
		}
	})

	if cmp.Diff(sampleArray, deserialized) != "" {
		t.Errorf("should be equal")
	}
}
