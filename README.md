# scale
### SCALE codec for go

Scale Codec written by golang, reference from [polkascan/py-scale-codec](https://github.com/polkascan/py-scale-codec).

---

# Interface
define a  interface ```Compact``` that means what we want to deal with. 
```go
type Compact interface {
	Encode() ([]byte, error)
	Decode(val []byte) (int, error)
	GetVal() interface{}
	GetType() PrimitiveType
	CloneNew() Compact
}
```

---


# Usage

Here are the examples for the types that we supported.

+ ##### Fix Number
use ```uint8``` as the example, we wrap the real val with a struct, the implement the interface. 
```go
// encode
c := FixU8{Val: 25}
res, err := c.Encode()


// decode
c := FixU8{}
res, err := c.Decode([]byte{25})
```

+ ##### Compact Number
compact number use the compact encode way which can save the space of memory.
```go
// encode
a := CompactU8{Val: uint8(25)}
res, err := a.Encode()

// decode
a := CompactU8{}
num, err := a.Decode([]byte{100})
```

+ ##### Bool
same as fix number, but the real value is a ```bool```.
```go
// encode
a := CompactBool{Val: true}
res, err := a.Encode()

// decode
a := CompactBool{}
num, err := a.Decode([]byte{1})
```

+ ##### String
```go
// encode
s := CompactString{
		Val: "set_hash",
	}
res, err := s.Encode()

// decode
s := &CompactString{}
s.Decode(Hex2Bytes("0x207365745f686173680c6b657914776f726c64"))
```

+ ##### Vec
in this type, we include ```NextList``` to define which type of value in the vector, we will also check the type while encode or decode. 
```go
// encode
v := &CompactVec{Val: []Compact{&FixU16{
Val: uint16(1),
}, &FixU16{
Val: uint16(64),
}}, NextList: []PrimitiveType{Uint16}}
res, err := v.Encode()

// decode
s := &CompactVec{NextList: []PrimitiveType{Uint16}}
_, err := s.Decode([]byte{8, 1, 0, 64, 0})
```

+ ##### Array

array is same to vector, but we add a attribute named ```Len``` to limit the number of value.
```go
// encode
v := &CompactArray{Val: []Compact{&FixU16{
Val: uint16(1),
}, &FixU16{
Val: uint16(64),
}}, NextList: []PrimitiveType{Uint16}, Len: 2}
res, err := v.Encode()

// decode
s := &CompactArray{NextList: []PrimitiveType{Uint16}, Len: 2}
_, err := s.Decode([]byte{1, 0, 64, 0})
```

+ ##### Enum
things to enum will be a little different. We do not encode or decode all types in the enum, but only choose one of it to operator. the attribute named ```Index``` tell us which element will be use.
```go
// for there is not enum type in golang, use rust for example
//  enum Demo {
//     First(i32),
//     Second(i64),
//  }
//

// encode
a := CompactEnum{Val: &FixU8{Val: 42}, Index: 0}
ans, err := a.Encode()

//decode
a := CompactEnum{Val: &FixU8{}, Index: 0}
ans, err := a.Decode([]byte{0, 42})
```

+ ##### Tuple
```go
// encode
a := CompactTuple{Val: []Compact{
			&CompactU32{Val: 3},
			&CompactBool{Val: false},
		}}
ans, err := a.Encode()

// decode
a := CompactTuple{Val: []Compact{
&CompactU32{},
&CompactBool{},
}}
_, err := a.Decode([]byte{12, 0})
```


