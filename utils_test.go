package scale

import (
	"github.com/stretchr/testify/assert"
	"math"
	"math/big"
	"testing"
)

func Test_convertToString(t *testing.T) {
	a, err := ConvertToString("123")
	assert.Nil(t, err)
	assert.Equal(t, "123", a.GetVal())

	_, err = ConvertToString(1)
	assert.NotNil(t, err)
}

func Test_convertToUint8(t *testing.T) {
	a, err := ConvertToUint8("123")
	assert.Nil(t, err)
	assert.Equal(t, uint8(123), a.GetVal())

	_, err = ConvertToUint8("test")
	assert.NotNil(t, err)

	a, err = ConvertToUint8(uint8(128))
	assert.Nil(t, err)
	assert.Equal(t, uint8(128), a.GetVal())

	a, err = ConvertToUint8(111)
	assert.Nil(t, err)
	assert.Equal(t, uint8(111), a.GetVal())

	_, err = ConvertToUint8(256)
	assert.NotNil(t, err)

	_, err = ConvertToUint8(uint16(23))
	assert.NotNil(t, err)
}

func Test_convertToInt8(t *testing.T) {
	a, err := ConvertToInt8("123")
	assert.Nil(t, err)
	assert.Equal(t, int8(123), a.GetVal())

	_, err = ConvertToInt8("test")
	assert.NotNil(t, err)

	a, err = ConvertToInt8(int8(123))
	assert.Nil(t, err)
	assert.Equal(t, int8(123), a.GetVal())

	_, err = ConvertToInt8(128)
	assert.NotNil(t, err)

	_, err = ConvertToInt8(256)
	assert.NotNil(t, err)

	_, err = ConvertToInt8(uint16(23))
	assert.NotNil(t, err)
}

func Test_convertToCompactU8(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		a, err := ConvertToCompactU8("123")
		assert.Nil(t, err)
		assert.Equal(t, uint8(123), a.GetVal())
	})
	t.Run("string+error", func(t *testing.T) {
		_, err := ConvertToCompactU8("test")
		assert.NotNil(t, err)
	})
	t.Run("uint8", func(t *testing.T) {
		a, err := ConvertToCompactU8(uint8(23))
		assert.Nil(t, err)
		assert.Equal(t, uint8(23), a.GetVal())
	})
	t.Run("int", func(t *testing.T) {
		a, err := ConvertToCompactU8(23)
		assert.Nil(t, err)
		assert.Equal(t, uint8(23), a.GetVal())
	})
	t.Run("int+error", func(t *testing.T) {
		_, err := ConvertToCompactU8(-1)
		assert.NotNil(t, err)
	})
	t.Run("err", func(t *testing.T) {
		_, err := ConvertToCompactU8(uint16(1))
		assert.NotNil(t, err)
	})
}

func Test_convertToUint16(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		a, err := ConvertToUint16("23456")
		assert.Nil(t, err)
		assert.Equal(t, uint16(23456), a.GetVal())
	})
	t.Run("string+error", func(t *testing.T) {
		_, err := ConvertToUint16("test")
		assert.NotNil(t, err)
	})
	t.Run("uint16", func(t *testing.T) {
		a, err := ConvertToUint16(uint16(25))
		assert.Nil(t, err)
		assert.Equal(t, uint16(25), a.GetVal())
	})
	t.Run("int", func(t *testing.T) {
		a, err := ConvertToUint16(24)
		assert.Nil(t, err)
		assert.Equal(t, uint16(24), a.GetVal())
	})
	t.Run("int+err", func(t *testing.T) {
		_, err := ConvertToUint16(-1)
		assert.NotNil(t, err)
	})
	t.Run("err", func(t *testing.T) {
		_, err := ConvertToUint16(uint8(1))
		assert.NotNil(t, err)
	})
}

func Test_convertToInt16(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		a, err := ConvertToInt16("123")
		assert.Nil(t, err)
		assert.Equal(t, int16(123), a.GetVal())
	})
	t.Run("string+err", func(t *testing.T) {
		_, err := ConvertToInt16("test")
		assert.NotNil(t, err)
	})
	t.Run("int16", func(t *testing.T) {
		a, err := ConvertToInt16(int16(30))
		assert.Nil(t, err)
		assert.Equal(t, int16(30), a.GetVal())
	})
	t.Run("int", func(t *testing.T) {
		a, err := ConvertToInt16(30)
		assert.Nil(t, err)
		assert.Equal(t, int16(30), a.GetVal())
	})
	t.Run("int+err", func(t *testing.T) {
		_, err := ConvertToInt16(math.MaxInt16 + 20)
		assert.NotNil(t, err)
	})
	t.Run("err", func(t *testing.T) {
		_, err := ConvertToInt16(uint8(20))
		assert.NotNil(t, err)
	})
}

func Test_convertToCompactU16(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		a, err := ConvertToCompactU16("123")
		assert.Nil(t, err)
		assert.Equal(t, uint16(123), a.GetVal())
	})
	t.Run("string+err", func(t *testing.T) {
		_, err := ConvertToCompactU16("test")
		assert.NotNil(t, err)
	})
	t.Run("uint16", func(t *testing.T) {
		a, err := ConvertToCompactU16(uint16(30))
		assert.Nil(t, err)
		assert.Equal(t, uint16(30), a.GetVal())
	})
	t.Run("int", func(t *testing.T) {
		a, err := ConvertToCompactU16(30)
		assert.Nil(t, err)
		assert.Equal(t, uint16(30), a.GetVal())
	})
	t.Run("int+err", func(t *testing.T) {
		_, err := ConvertToCompactU16(math.MaxUint16 + 2)
		assert.NotNil(t, err)
	})
	t.Run("err", func(t *testing.T) {
		_, err := ConvertToCompactU16(uint32(1))
		assert.NotNil(t, err)
	})
}

func Test_convertToUint32(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		a, err := ConvertToUint32("12345")
		assert.Nil(t, err)
		assert.Equal(t, uint32(12345), a.GetVal())
	})
	t.Run("string+err", func(t *testing.T) {
		_, err := ConvertToUint32("test")
		assert.NotNil(t, err)
	})
	t.Run("uint32", func(t *testing.T) {
		a, err := ConvertToUint32(uint32(45678))
		assert.Nil(t, err)
		assert.Equal(t, uint32(45678), a.GetVal())
	})
	t.Run("int", func(t *testing.T) {
		a, err := ConvertToUint32(1234)
		assert.Nil(t, err)
		assert.Equal(t, uint32(1234), a.GetVal())
	})
	t.Run("int+err", func(t *testing.T) {
		_, err := ConvertToUint32(math.MaxInt64)
		assert.NotNil(t, err)
	})
	t.Run("err", func(t *testing.T) {
		_, err := ConvertToUint32(uint8(1))
		assert.NotNil(t, err)
	})
}

func Test_convertToInt32(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		a, err := ConvertToInt32("123")
		assert.Nil(t, err)
		assert.Equal(t, int32(123), a.GetVal())
	})
	t.Run("string+err", func(t *testing.T) {
		_, err := ConvertToInt32("test")
		assert.NotNil(t, err)
	})
	t.Run("int32", func(t *testing.T) {
		a, err := ConvertToInt32(int32(123))
		assert.Nil(t, err)
		assert.Equal(t, int32(123), a.GetVal())
	})
	t.Run("int", func(t *testing.T) {
		a, err := ConvertToInt32(123)
		assert.Nil(t, err)
		assert.Equal(t, int32(123), a.GetVal())
	})
	t.Run("int+err", func(t *testing.T) {
		_, err := ConvertToInt32(math.MaxInt32 + 3)
		assert.NotNil(t, err)
	})
	t.Run("err", func(t *testing.T) {
		_, err := ConvertToInt32(uint16(1))
		assert.NotNil(t, err)
	})
}

func Test_convertToCompactU32(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		a, err := ConvertToCompactU32("123")
		assert.Nil(t, err)
		assert.Equal(t, uint32(123), a.GetVal())
	})
	t.Run("string+err", func(t *testing.T) {
		_, err := ConvertToCompactU32("test")
		assert.NotNil(t, err)
	})
	t.Run("uint32", func(t *testing.T) {
		a, err := ConvertToCompactU32(uint32(123))
		assert.Nil(t, err)
		assert.Equal(t, uint32(123), a.GetVal())
	})
	t.Run("int", func(t *testing.T) {
		a, err := ConvertToCompactU32(123)
		assert.Nil(t, err)
		assert.Equal(t, uint32(123), a.GetVal())
	})
	t.Run("int+err", func(t *testing.T) {
		_, err := ConvertToCompactU32(math.MaxUint32 + 2)
		assert.NotNil(t, err)
	})
	t.Run("err", func(t *testing.T) {
		_, err := ConvertToCompactU32(int16(32))
		assert.NotNil(t, err)
	})
}

func Test_convertToUint64(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		a, err := ConvertToUint64("1234")
		assert.Nil(t, err)
		assert.Equal(t, uint64(1234), a.GetVal())
	})
	t.Run("string+err", func(t *testing.T) {
		_, err := ConvertToUint64("test")
		assert.NotNil(t, err)
	})
	t.Run("uint64", func(t *testing.T) {
		a, err := ConvertToUint64(uint64(123))
		assert.Nil(t, err)
		assert.Equal(t, uint64(123), a.GetVal())
	})
	t.Run("int", func(t *testing.T) {
		a, err := ConvertToUint64(123)
		assert.Nil(t, err)
		assert.Equal(t, uint64(123), a.GetVal())
	})
	t.Run("int+err", func(t *testing.T) {
		_, err := ConvertToUint64(-1)
		assert.NotNil(t, err)
	})
	t.Run("err", func(t *testing.T) {
		_, err := ConvertToUint64(int16(-1))
		assert.NotNil(t, err)
	})
}

func Test_convertToInt64(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		a, err := ConvertToInt64("1234")
		assert.Nil(t, err)
		assert.Equal(t, int64(1234), a.GetVal())
	})
	t.Run("string+err", func(t *testing.T) {
		_, err := ConvertToInt64("test")
		assert.NotNil(t, err)
	})
	t.Run("int64", func(t *testing.T) {
		a, err := ConvertToInt64(int64(-1))
		assert.Nil(t, err)
		assert.Equal(t, int64(-1), a.GetVal())
	})
	t.Run("int", func(t *testing.T) {
		a, err := ConvertToInt64(123)
		assert.Nil(t, err)
		assert.Equal(t, int64(123), a.GetVal())
	})
	t.Run("err", func(t *testing.T) {
		_, err := ConvertToInt64(int16(-1))
		assert.NotNil(t, err)
	})
}

func Test_convertToCompactU64(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		a, err := ConvertToCompactU64("1234")
		assert.Nil(t, err)
		assert.Equal(t, uint64(1234), a.GetVal())
	})
	t.Run("string+err", func(t *testing.T) {
		_, err := ConvertToCompactU64("test")
		assert.NotNil(t, err)
	})
	t.Run("uint64", func(t *testing.T) {
		a, err := ConvertToCompactU64(uint64(1))
		assert.Nil(t, err)
		assert.Equal(t, uint64(1), a.GetVal())
	})
	t.Run("int", func(t *testing.T) {
		a, err := ConvertToCompactU64(123)
		assert.Nil(t, err)
		assert.Equal(t, uint64(123), a.GetVal())
	})
	t.Run("int+err", func(t *testing.T) {
		_, err := ConvertToCompactU64(-1)
		assert.NotNil(t, err)
	})
	t.Run("err", func(t *testing.T) {
		_, err := ConvertToCompactU64(uint16(1))
		assert.NotNil(t, err)
	})
}

func Test_convertToInt128(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		a, err := ConvertToInt128("1234")
		assert.Nil(t, err)
		assert.Equal(t, uint64(1234), a.GetVal().(*big.Int).Uint64())
	})
	t.Run("string+err", func(t *testing.T) {
		_, err := ConvertToInt128("test")
		assert.NotNil(t, err)
	})
	t.Run("int128", func(t *testing.T) {
		a, err := ConvertToInt128(new(big.Int).SetInt64(1234))
		assert.Nil(t, err)
		assert.Equal(t, uint64(1234), a.GetVal().(*big.Int).Uint64())
	})
	t.Run("int", func(t *testing.T) {
		a, err := ConvertToInt128(1234)
		assert.Nil(t, err)
		assert.Equal(t, uint64(1234), a.GetVal().(*big.Int).Uint64())
	})
	t.Run("err", func(t *testing.T) {
		_, err := ConvertToInt128(int16(1234))
		assert.NotNil(t, err)
	})
}

func Test_convertToUint128(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		a, err := ConvertToUint128("1234")
		assert.Nil(t, err)
		assert.Equal(t, uint64(1234), a.GetVal().(*big.Int).Uint64())
	})
	t.Run("string+err", func(t *testing.T) {
		_, err := ConvertToUint128("test")
		assert.NotNil(t, err)
	})
	t.Run("int128", func(t *testing.T) {
		a, err := ConvertToUint128(new(big.Int).SetInt64(1234))
		assert.Nil(t, err)
		assert.Equal(t, uint64(1234), a.GetVal().(*big.Int).Uint64())
	})
	t.Run("int", func(t *testing.T) {
		a, err := ConvertToUint128(1234)
		assert.Nil(t, err)
		assert.Equal(t, uint64(1234), a.GetVal().(*big.Int).Uint64())
	})
	t.Run("err", func(t *testing.T) {
		_, err := ConvertToUint128(int16(1234))
		assert.NotNil(t, err)
	})
}

func Test_convertToCompactU128(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		a, err := ConvertToCompactU128("1234")
		assert.Nil(t, err)
		assert.Equal(t, uint64(1234), a.GetVal().(*big.Int).Uint64())
	})
	t.Run("string+err", func(t *testing.T) {
		_, err := ConvertToCompactU128("test")
		assert.NotNil(t, err)
	})
	t.Run("int128", func(t *testing.T) {
		a, err := ConvertToCompactU128(new(big.Int).SetInt64(1234))
		assert.Nil(t, err)
		assert.Equal(t, uint64(1234), a.GetVal().(*big.Int).Uint64())
	})
	t.Run("int", func(t *testing.T) {
		a, err := ConvertToCompactU128(1234)
		assert.Nil(t, err)
		assert.Equal(t, uint64(1234), a.GetVal().(*big.Int).Uint64())
	})
	t.Run("err", func(t *testing.T) {
		_, err := ConvertToCompactU128(int16(1234))
		assert.NotNil(t, err)
	})
}

func Test_convertToBool(t *testing.T) {
	t.Run("bool", func(t *testing.T) {
		a, err := ConvertToBool(true)
		assert.Nil(t, err)
		assert.Equal(t, true, a.GetVal())
	})
	t.Run("err", func(t *testing.T) {
		_, err := ConvertToBool("test")
		assert.NotNil(t, err)
	})
}

func Test_getCompactIntEncodeMode(t *testing.T) {
	t.Run("0b00", func(t *testing.T) {
		assert.Equal(t, 0, getCompactIntEncodeMode(42))
	})
	t.Run("0b01", func(t *testing.T) {
		assert.Equal(t, 1, getCompactIntEncodeMode(69))
	})
	t.Run("0b10", func(t *testing.T) {
		assert.Equal(t, 2, getCompactIntEncodeMode(1<<14))
	})
	t.Run("0b11", func(t *testing.T) {
		assert.Equal(t, 3, getCompactIntEncodeMode(1<<30))
	})
}
