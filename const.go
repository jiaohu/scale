package scale

const (
	// MoveBit for the highest six bit is has special means, in order to get origin num, must move right 2bit
	// also encode a number must move left 2bit
	MoveBit = 2

	// SingleMode the higher six bit is LE for value, which is 0 to 63
	SingleMode = 0b00
	MaxSingle  = 63
	// DoubleMode the higher six bit and next byte is LE for value, which is 64 to 1<<14 - 1
	DoubleMode = 0b01
	MaxDouble  = 1<<14 - 1
	// FourByteMode the higher six bit and next three bytes is LE for value, which is 1<<14 -1 to 1<<30-1
	FourByteMode = 0b10
	MaxFour      = 1<<30 - 1
	// BigIntMode the highest six bit represent the bytes num in the next, which less than 4.
	// then, the next bytes contain the LE for value. the highest one can not be 0, which is
	// 1<<30-1 to 1<<536-1
	BigIntMode = 0b11
)

type PrimitiveType int

const (
	None PrimitiveType = iota
	String
	Struct
	Uint8
	Int8
	CompactUInt8
	Uint16
	Int16
	CompactUint16
	Uint32
	Int32
	CompactUint32
	Uint64
	Int64
	CompactUint64
	BigInt
	BigUint
	CompactBigInt
	Bool
	Vec
	Primitive
	Array
	Tuple
	Enum
)
