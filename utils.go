package scale

import (
	"encoding/hex"
	"errors"
	"math"
	"math/big"
	"strconv"
	"strings"
)

func reverse(input []byte) (output []byte) {
	for i := len(input) - 1; i >= 0; i-- {
		output = append(output, input[i])
	}
	return
}

func negchange(input []byte) []byte {
	for i, t := range input {
		input[i] = math.MaxUint8 - t
	}
	incr := 1
	for i := len(input) - 1; i >= 0; i-- {
		temp := int(input[i]) + incr
		if temp > math.MaxUint8 {
			input[i] = 0
		} else {
			input[i] = byte(temp)
			break
		}
	}
	return input
}

func ConvertToString(param interface{}) (*CompactString, error) {
	if val, ok := param.(string); ok {
		return &CompactString{Val: val}, nil
	} else {
		return nil, errors.New("param not string")
	}
}

func ConvertToUint8(param interface{}) (*FixU8, error) {
	if val, ok := param.(uint8); ok {
		return &FixU8{Val: val}, nil
	} else if val2, ok2 := param.(string); ok2 {
		s, err := strconv.ParseUint(val2, 10, 8)
		if err != nil {
			return nil, err
		}
		return &FixU8{Val: uint8(s)}, nil
	} else if val3, ok3 := param.(int); ok3 {
		if val3 >= 0 && val3 <= math.MaxUint8 {
			return &FixU8{Val: uint8(val3)}, nil
		} else {
			return nil, errors.New("out of uint8 range")
		}
	} else {
		return nil, errors.New("param not uint8")
	}
}

func ConvertToInt8(param interface{}) (*FixI8, error) {
	if val, ok := param.(int8); ok {
		return &FixI8{Val: val}, nil
	} else if val2, ok2 := param.(string); ok2 {
		s, err := strconv.ParseInt(val2, 10, 8)
		if err != nil {
			return nil, err
		}
		return &FixI8{Val: int8(s)}, nil
	} else if val3, ok3 := param.(int); ok3 {
		if val3 >= math.MinInt8 && val3 <= math.MaxInt8 {
			return &FixI8{Val: int8(val3)}, nil
		} else {
			return nil, errors.New("out of int8 range")
		}
	} else {
		return nil, errors.New("param not int8")
	}
}

func ConvertToCompactU8(param interface{}) (*CompactU8, error) {
	if val, ok := param.(uint8); ok {
		return &CompactU8{Val: val}, nil
	} else if val2, ok2 := param.(string); ok2 {
		s, err := strconv.ParseUint(val2, 10, 8)
		if err != nil {
			return nil, err
		}
		return &CompactU8{Val: uint8(s)}, nil
	} else if val3, ok3 := param.(int); ok3 {
		if val3 >= 0 && val3 <= math.MaxUint8 {
			return &CompactU8{Val: uint8(val3)}, nil
		} else {
			return nil, errors.New("out of uint8 range")
		}
	} else {
		return nil, errors.New("param not compact uint8")
	}
}

func ConvertToUint16(param interface{}) (*FixU16, error) {
	if val, ok := param.(uint16); ok {
		return &FixU16{Val: val}, nil
	} else if val2, ok2 := param.(string); ok2 {
		s, err := strconv.ParseUint(val2, 10, 16)
		if err != nil {
			return nil, err
		}
		return &FixU16{Val: uint16(s)}, nil
	} else if val3, ok3 := param.(int); ok3 {
		if val3 >= 0 && val3 <= math.MaxUint16 {
			return &FixU16{Val: uint16(val3)}, nil
		} else {
			return nil, errors.New("out of uint16 range")
		}
	} else {
		return nil, errors.New("param not uint16")
	}
}

func ConvertToInt16(param interface{}) (*FixI16, error) {
	if val, ok := param.(int16); ok {
		return &FixI16{Val: val}, nil
	} else if val2, ok2 := param.(string); ok2 {
		s, err := strconv.ParseInt(val2, 10, 16)
		if err != nil {
			return nil, err
		}
		return &FixI16{Val: int16(s)}, nil
	} else if val3, ok3 := param.(int); ok3 {
		if val3 >= math.MinInt16 && val3 <= math.MaxInt16 {
			return &FixI16{Val: int16(val3)}, nil
		} else {
			return nil, errors.New("out of int16 range")
		}
	} else {
		return nil, errors.New("param not int16")
	}
}

func ConvertToCompactU16(param interface{}) (*CompactU16, error) {
	if val, ok := param.(uint16); ok {
		return &CompactU16{Val: val}, nil
	} else if val2, ok2 := param.(string); ok2 {
		s, err := strconv.ParseUint(val2, 10, 16)
		if err != nil {
			return nil, err
		}
		return &CompactU16{Val: uint16(s)}, nil
	} else if val3, ok3 := param.(int); ok3 {
		if val3 >= 0 && val3 <= math.MaxUint16 {
			return &CompactU16{Val: uint16(val3)}, nil
		} else {
			return nil, errors.New("out of uint16 range")
		}
	} else {
		return nil, errors.New("param not compact uint16")
	}
}

func ConvertToUint32(param interface{}) (*FixU32, error) {
	if val, ok := param.(uint32); ok {
		return &FixU32{Val: val}, nil
	} else if val2, ok2 := param.(string); ok2 {
		s, err := strconv.ParseUint(val2, 10, 32)
		if err != nil {
			return nil, err
		}
		return &FixU32{Val: uint32(s)}, nil
	} else if val3, ok3 := param.(int); ok3 {
		if val3 >= 0 && val3 <= math.MaxUint32 {
			return &FixU32{Val: uint32(val3)}, nil
		} else {
			return nil, errors.New("out of uint32 range")
		}
	} else {
		return nil, errors.New("param not uint32")
	}
}

func ConvertToInt32(param interface{}) (*FixI32, error) {
	if val, ok := param.(int32); ok {
		return &FixI32{Val: val}, nil
	} else if val2, ok2 := param.(string); ok2 {
		s, err := strconv.ParseInt(val2, 10, 32)
		if err != nil {
			return nil, err
		}
		return &FixI32{Val: int32(s)}, nil
	} else if val3, ok3 := param.(int); ok3 {
		if val3 >= math.MinInt32 && val3 <= math.MaxInt32 {
			return &FixI32{Val: int32(val3)}, nil
		} else {
			return nil, errors.New("out of int32 range")
		}
	} else {
		return nil, errors.New("param not int32")
	}
}

func ConvertToCompactU32(param interface{}) (*CompactU32, error) {
	if val, ok := param.(uint32); ok {
		return &CompactU32{Val: val}, nil
	} else if val2, ok2 := param.(string); ok2 {
		s, err := strconv.ParseUint(val2, 10, 32)
		if err != nil {
			return nil, err
		}
		return &CompactU32{Val: uint32(s)}, nil
	} else if val3, ok3 := param.(int); ok3 {
		if val3 >= 0 && val3 <= math.MaxUint32 {
			return &CompactU32{Val: uint32(val3)}, nil
		} else {
			return nil, errors.New("out of uint32 range")
		}
	} else {
		return nil, errors.New("param not compact uint32")
	}
}

func ConvertToUint64(param interface{}) (*FixU64, error) {
	if val, ok := param.(uint64); ok {
		return &FixU64{Val: val}, nil
	} else if val2, ok2 := param.(string); ok2 {
		s, err := strconv.ParseUint(val2, 10, 64)
		if err != nil {
			return nil, err
		}
		return &FixU64{Val: s}, nil
	} else if val3, ok3 := param.(int); ok3 {
		if val3 >= 0 && uint64(val3) <= math.MaxUint64 {
			return &FixU64{Val: uint64(val3)}, nil
		} else {
			return nil, errors.New("out of uint64 range")
		}
	} else {
		return nil, errors.New("param not uint64")
	}
}

func ConvertToInt64(param interface{}) (*FixI64, error) {
	if val, ok := param.(int64); ok {
		return &FixI64{Val: val}, nil
	} else if val2, ok2 := param.(string); ok2 {
		s, err := strconv.ParseInt(val2, 10, 64)
		if err != nil {
			return nil, err
		}
		return &FixI64{Val: s}, nil
	} else if val3, ok3 := param.(int); ok3 {
		if val3 >= math.MinInt64 && val3 <= math.MaxInt64 {
			return &FixI64{Val: int64(val3)}, nil
		} else {
			return nil, errors.New("out of int64 range")
		}
	} else {
		return nil, errors.New("param not int64")
	}
}

func ConvertToCompactU64(param interface{}) (*CompactU64, error) {
	if val, ok := param.(uint64); ok {
		return &CompactU64{Val: val}, nil
	} else if val2, ok2 := param.(string); ok2 {
		s, err := strconv.ParseUint(val2, 10, 64)
		if err != nil {
			return nil, err
		}
		return &CompactU64{Val: s}, nil
	} else if val3, ok3 := param.(int); ok3 {
		if val3 >= 0 && uint64(val3) <= math.MaxUint64 {
			return &CompactU64{Val: uint64(val3)}, nil
		} else {
			return nil, errors.New("out of uint64 range")
		}
	} else {
		return nil, errors.New("param not compact uint64")
	}
}

func ConvertToInt128(param interface{}) (*FixI128, error) {
	if val, ok := param.(*big.Int); ok {
		return &FixI128{Val: val}, nil
	} else if val2, ok2 := param.(string); ok2 {
		t := new(big.Int)
		_, success := t.SetString(val2, 10)
		if !success {
			return nil, errors.New("invalid param")
		}
		return &FixI128{Val: t}, nil
	} else if val3, ok3 := param.(int); ok3 {
		t := new(big.Int)
		t.SetInt64(int64(val3))
		return &FixI128{Val: t}, nil
	} else {
		return nil, errors.New("param not i128")
	}
}

func ConvertToUint128(param interface{}) (*FixU128, error) {
	if val, ok := param.(*big.Int); ok {
		return &FixU128{Val: val}, nil
	} else if val2, ok2 := param.(string); ok2 {
		t := new(big.Int)
		_, success := t.SetString(val2, 10)
		if !success {
			return nil, errors.New("invalid param")
		}
		return &FixU128{Val: t}, nil
	} else if val3, ok3 := param.(int); ok3 {
		t := new(big.Int)
		t.SetUint64(uint64(val3))
		return &FixU128{Val: t}, nil
	} else {
		return nil, errors.New("param not u128")
	}
}

func ConvertToCompactU128(param interface{}) (*CompactU128, error) {
	if val, ok := param.(*big.Int); ok {
		return &CompactU128{Val: val}, nil
	} else if val2, ok2 := param.(string); ok2 {
		t := new(big.Int)
		_, success := t.SetString(val2, 10)
		if !success {
			return nil, errors.New("invalid param")
		}
		return &CompactU128{Val: t}, nil
	} else if val3, ok3 := param.(int); ok3 {
		t := new(big.Int)
		t.SetUint64(uint64(val3))
		return &CompactU128{Val: t}, nil
	} else {
		return nil, errors.New("param not compact u128")
	}
}

func ConvertToBool(param interface{}) (*CompactBool, error) {
	if val, ok := param.(bool); ok {
		return &CompactBool{Val: val}, nil
	} else {
		return nil, errors.New("param not bool")
	}
}

func getCompactIntEncodeMode(a int) int {
	if a >= 0 && a <= MaxSingle {
		return SingleMode
	} else if a <= MaxDouble {
		return DoubleMode
	} else if a <= MaxFour {
		return FourByteMode
	} else {
		return BigIntMode
	}
}

// GetCompactValue get Compact's real value
func GetCompactValue(val Compact) interface{} {
	switch val.GetType() {
	case Array, Vec, Struct, Tuple:
		var values []interface{}
		for _, v1 := range val.GetVal().([]Compact) {
			values = append(values, GetCompactValue(v1))
		}
		return values
	default:
		return val.GetVal()
	}
}

func Hex2Bytes(str string) []byte {
	if strings.HasPrefix(str, "0x") {
		str = str[2:]
	}
	h, _ := hex.DecodeString(str)
	return h
}

func Bytes2Hex(d []byte) string {
	return hex.EncodeToString(d)
}
