package scale

import (
	"bytes"
	"errors"
)

type CompactArray struct {
	Val     []Compact
	Len     int
	EleType PrimitiveType
}

func (c *CompactArray) Encode() ([]byte, error) {
	if len(c.Val) != c.Len {
		return nil, errors.New("wrong array length")
	}
	if c.EleType == 0 {
		return nil, errors.New("must point element type for array")
	}
	var buf bytes.Buffer
	for _, v := range c.Val {
		if v.GetType() != c.EleType {
			return nil, errors.New("invalid array")
		}
		res, err := v.Encode()
		if err != nil {
			return nil, err
		}
		buf.Write(res)
	}
	return buf.Bytes(), nil
}

func (c *CompactArray) Decode(value []byte) (int, error) {
	offset := 0
	for i := 0; i < c.Len; i++ {
		temp, err := c.getNextCompact(i)
		if err != nil {
			return 0, err
		}
		tempOffset, err := temp.Decode(value)
		if err != nil {
			return 0, err
		}
		if len(c.Val)-1 < i {
			c.Val = append(c.Val, temp)
		} else {
			c.Val[i] = temp
		}
		offset += tempOffset
		value = value[tempOffset:]
	}
	return offset, nil
}

func (c *CompactArray) GetVal() interface{} {
	return c.Val
}

func (c *CompactArray) GetType() PrimitiveType {
	return Array
}

func (c *CompactArray) getNextCompact(i int) (Compact, error) {
	switch c.EleType {
	case String:
		return &CompactString{}, nil
	case Uint8:
		return &FixU8{}, nil
	case Int8:
		return &FixI8{}, nil
	case CompactUInt8:
		return &CompactU8{}, nil
	case Uint16:
		return &FixU16{}, nil
	case Int16:
		return &FixI16{}, nil
	case CompactUint16:
		return &CompactU16{}, nil
	case Uint32:
		return &FixU32{}, nil
	case Int32:
		return &FixI32{}, nil
	case CompactUint32:
		return &CompactU32{}, nil
	case Uint64:
		return &FixU64{}, nil
	case Int64:
		return &FixI64{}, nil
	case CompactUint64:
		return &CompactU64{}, nil
	case BigInt:
		return &FixI128{}, nil
	case BigUint:
		return &FixU128{}, nil
	case CompactBigInt:
		return &CompactU128{}, nil
	case Bool:
		return &CompactBool{}, nil
	case Vec:
		return c.Val[i].(*CompactVec), nil
	case Struct:
		return c.Val[i].(*CompactStruct), nil
	case Array:
		return c.Val[i].(*CompactArray), nil
	default:
		return nil, errors.New("not supported type")
	}
}

func (c *CompactArray) CloneNew() Compact {
	temp := &CompactArray{
		Len: c.Len,
	}
	temp.EleType = c.EleType
	for _, v := range c.Val {
		temp.Val = append(temp.Val, v.CloneNew())
	}
	return temp
}
