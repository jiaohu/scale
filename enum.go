package scale

import "bytes"

type CompactEnum struct {
	Val   Compact
	Index byte
}

func (c *CompactEnum) Encode() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte(c.Index)
	if c.Val == nil {
		return buf.Bytes(), nil
	}
	res, err := c.Val.Encode()
	if err != nil {
		return nil, err
	}
	buf.Write(res)
	return buf.Bytes(), nil
}

func (c *CompactEnum) Decode(value []byte) (int, error) {
	var offset int
	if len(value) == 0 {
		return 0, nil
	}
	c.Index = value[0]
	offset += 1
	value = value[1:]
	if c.Val == nil {
		return offset, nil
	}
	tempOffset, err := c.Val.Decode(value)
	if err != nil {
		return 0, err
	}
	offset += tempOffset
	value = value[tempOffset:]
	return offset, nil
}

func (c *CompactEnum) GetVal() interface{} {
	if c.Val == nil {
		return nil
	}
	return c.Val.GetVal()
}

func (c *CompactEnum) GetType() PrimitiveType {
	return Enum
}

func (c *CompactEnum) CloneNew() Compact {
	temp := &CompactEnum{}
	temp.Val = c.Val.CloneNew()
	temp.Index = c.Index
	return temp
}
