package types

import "fmt"

type ValueType byte

const (
	TypeNone ValueType = iota
	TypeAny
	TypeI32
	TypeI64
	TypeF32
	TypeF64
	TypeV128
	TypeAnyRef
	TypeAnyFunc
	TypeNullRef

	TypeMax = TypeNullRef - TypeNone
)

const (
	ErrInvalidValueType = "Invalid valud type"
)

func (vt ValueType) String() string {
	switch vt {
	case TypeAny:
		return "type:any"
	case TypeI32:
		return "type:i32"
	case TypeI64:
		return "type:i64"
	case TypeF32:
		return "type:f32"
	case TypeF64:
		return "type:f64"
	case TypeV128:
		return "type:v128"
	case TypeAnyRef:
		return "type:anyRef"
	case TypeAnyFunc:
		return "type:anyFunc"
	case TypeNullRef:
		return "type:nullRef"
	}
	return "type:none"
}

func DecodeValueType(vt int32) (ValueType, error) {
	switch vt {
	case -1:
		return TypeI32, nil;
	case -2:
		return TypeI64, nil;
	case -3:
		return TypeF32, nil;
	case -4:
		return TypeF64, nil;
	case -5:
		return TypeV128, nil;
	case -16:
		return TypeAnyFunc, nil;
	case -17:
		return TypeAnyRef, nil;
	}
	return TypeNone, fmt.Errorf(ErrInvalidValueType)
}
