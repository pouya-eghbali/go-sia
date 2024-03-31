package go_sia__test

import (
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

type Sample struct {
	Name      string
	Age       uint8
	IsSingle  bool
	Tags      []string
	BirthDate *big.Int
}

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
