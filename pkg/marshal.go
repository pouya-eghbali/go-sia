package sia

import (
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
)

func Marshal(v any) ([]byte, error) {
	s := New()
	t := reflect.TypeOf(v)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := reflect.ValueOf(v).Field(i).Interface()

		fieldType := field.Type.Kind()
		if fieldType == reflect.Ptr {
			fieldType = field.Type.Elem().Kind()
		}

		switch fieldType {
		case reflect.Uint8:
			s.AddUInt8(fieldValue.(uint8))
		case reflect.Uint16:
			s.AddUInt16(fieldValue.(uint16))
		case reflect.Uint32:
			s.AddUInt32(fieldValue.(uint32))
		case reflect.Uint64:
			s.AddUInt64(fieldValue.(uint64))

		case reflect.Int8:
			s.AddInt64(fieldValue.(int64))
		case reflect.Int16:
			s.AddInt16(fieldValue.(int16))
		case reflect.Int32:
			s.AddInt32(fieldValue.(int32))
		case reflect.Int64:
			s.AddInt64(fieldValue.(int64))

		case reflect.Bool:
			s.AddBool(fieldValue.(bool))

		case reflect.Slice:
			arrayByte, err := json.Marshal(fieldValue)
			if err != nil {
				panic(err)
			}
			s.AddByteArray16(arrayByte)
		case reflect.Array:
			arrayByte, err := json.Marshal(fieldValue)
			if err != nil {
				panic(err)
			}
			s.AddByteArray16(arrayByte)

		case reflect.TypeOf(big.Int{}).Kind():
			s.AddBigInt(fieldValue.(*big.Int))
		case reflect.String:
			if len(fieldValue.(string)) <= 8 {
				s.AddString8(fieldValue.(string))
			} else if len(fieldValue.(string)) <= 16 {
				s.AddString16(fieldValue.(string))
			} else if len(fieldValue.(string)) <= 32 {
				s.AddString32(fieldValue.(string))
			} else if len(fieldValue.(string)) <= 64 {
				s.AddString64(fieldValue.(string))
			}
		default:
			return nil, fmt.Errorf("type %s is not supported", fieldType.String())
		}
	}

	return s.Byte(), nil
}

func Unmarshal(data []byte, v any) error {
	s := NewFromByte(data)
	t := reflect.TypeOf(v).Elem()
	rv := reflect.ValueOf(v).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := rv.Field(i)

		fieldValue := rv.Field(i).Interface()

		fieldType := field.Type.Kind()
		if fieldType == reflect.Ptr {
			fieldType = field.Type.Elem().Kind()
		}

		switch fieldType {
		case reflect.Uint8:
			value.Set(reflect.ValueOf(s.ReadUInt8()))
		case reflect.Uint16:
			value.Set(reflect.ValueOf(s.ReadUInt16()))
		case reflect.Uint32:
			value.Set(reflect.ValueOf(s.ReadUInt32()))
		case reflect.Uint64:
			value.Set(reflect.ValueOf(s.ReadUInt64()))

		case reflect.Int8:
			value.Set(reflect.ValueOf(s.ReadInt64()))
		case reflect.Int16:
			value.Set(reflect.ValueOf(s.ReadInt16()))
		case reflect.Int32:
			value.Set(reflect.ValueOf(s.ReadInt32()))
		case reflect.Int64:
			value.Set(reflect.ValueOf(s.ReadInt64()))

		case reflect.Bool:
			value.Set(reflect.ValueOf(s.ReadBool()))

		case reflect.Slice:
			arrayByte := s.ReadByteArray16()
			err := json.Unmarshal(arrayByte, value.Addr().Interface())
			if err != nil {
				panic(err)
			}
		case reflect.Array:
			arrayByte := s.ReadByteArray16()
			err := json.Unmarshal(arrayByte, value.Addr().Interface())
			if err != nil {
				panic(err)
			}
		case reflect.TypeOf(big.Int{}).Kind():
			value.Set(reflect.ValueOf(s.ReadBigInt()))
		case reflect.String:
			if len(fieldValue.(string)) <= 8 {
				value.SetString(s.ReadString8())
			} else if len(fieldValue.(string)) <= 16 {
				value.SetString(s.ReadString16())
			} else if len(fieldValue.(string)) <= 32 {
				value.SetString(s.ReadString32())
			} else if len(fieldValue.(string)) <= 64 {
				value.SetString(s.ReadString64())
			}

		default:
			return fmt.Errorf("type %s is not supported", fieldType.String())
		}
	}

	return nil
}
