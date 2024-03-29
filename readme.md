# Sia

[![GoDev](https://img.shields.io/static/v1?label=godev&message=reference&color=00add8)][godev]
[![Build Status](https://github.com/pouya-eghbali/go-sia/actions/workflows/test.yml/badge.svg?branch=master)][actions]

Sia - Binary serialisation and deserialisation with built-in compression. You can consider Sia a strongly typed, statically typed domain specific binary language for constructing data. Sia preserves data types and supports custom ones.

## Install

`
go get github.com/pouya-eghbali/go-sia
`

## Basic Usage

To serialize multiple values, first create a sia object and then you can add values in order. Note that the order of adding values should be considered when you want to read them again.

Serializing:
```go
rawByte := sia.New().
    AddUInt16(1234).
    AddString64("think simple, do simple!").
    Byte()
```

Deserializing:
```go
deserialized := sia.NewFromByte(rawByte)
gotSampleUint16 := deserialized.ReadUInt16() // 1234
gotSampleString := deserialized.ReadString64() // think simple, do simple!
```

## Serialize Struct

If you want to serialize an object at once, you can use the Marshal and Unmarshal. Note that the order of structs is important here. Sia don't save the name of fields, so deserializing objects will be on the order of fields.

Serializing:
```go
sampleStruct := Sample{
    Name:      "test",
    Age:       18,
    IsSingle:  true,
    Tags:      []string{"t1", "t2"},
    BirthDate: big.NewInt(12345678987654321),
}

rawByte, err := sia.Marshal(sampleStruct)
```

Deserializing:
```go
sample := Sample{}
err = sia.Unmarshal(rawByte, &gotSample)
```

Note that sia can't handle serializing of arrays, so it will fall back to JSON marshal about them.