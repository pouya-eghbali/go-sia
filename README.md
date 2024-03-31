# Sia

[![Build Status](https://github.com/logicalangel/go-sia/actions/workflows/test.yml/badge.svg?branch=master)][actions]

Sia - Binary serialisation and deserialisation with built-in compression. You can consider Sia a strongly typed, statically typed domain specific binary language for constructing data. Sia preserves data types and supports custom ones.

[actions]: https://github.com/logicalangel/go-sia

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

Note that sia can't handle serializing of arrays, so it will fall back to JSON marshal about them.