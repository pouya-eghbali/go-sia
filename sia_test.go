package go_sia__test

import (
	"github.com/google/go-cmp/cmp"
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
		Byte()

	deserialized := sia.NewFromByte(rawByte)
	gotSampleUint16 := deserialized.ReadUInt16()
	assert.Equal(t, sampleUint16, gotSampleUint16, "should be equal")
	gotSampleString := deserialized.ReadString64()
	assert.Equal(t, sampleString, gotSampleString, "should be equal")
}

func TestMarshal(t *testing.T) {
	sampleStruct := Sample{
		Name:      "test",
		Age:       18,
		IsSingle:  true,
		Tags:      []string{"t1", "t2"},
		BirthDate: big.NewInt(12345678987654321),
	}

	rawByte, err := sia.Marshal(sampleStruct)
	assert.Nil(t, err)

	gotSample := Sample{}
	err = sia.Unmarshal(rawByte, &gotSample)
	assert.Nil(t, err)

	assert.Empty(t, cmp.Diff(sampleStruct, gotSample, cmp.AllowUnexported(big.Int{})))
}
