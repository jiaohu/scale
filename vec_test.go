package scale

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompactVec_Encode(t *testing.T) {
	v := &CompactVec{Val: []Compact{&FixU16{
		Val: uint16(1),
	}, &FixU16{
		Val: uint16(64),
	}}, EleType: Uint16}
	res, err := v.Encode()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, []byte{8, 1, 0, 64, 0}, res)
}

func TestCompactVec_Decode(t *testing.T) {
	s := &CompactVec{EleType: Uint16}
	_, err := s.Decode([]byte{8, 1, 0, 64, 0})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, uint16(1), s.Val[0].GetVal())
	assert.Equal(t, uint16(64), s.Val[1].GetVal())
}
