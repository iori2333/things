package json

import (
	"strings"
)

type Value interface {
	Int() (int, bool)
	Float() (float64, bool)
	Bool() (bool, bool)
	String() (string, bool)
	Array() (Array, bool)
	Object() (Object, bool)
	Type() Type
	Clone() Value
}

type Number float64

func (n Number) Int() (int, bool) {
	return int(n), true
}

func (n Number) Float() (float64, bool) {
	return float64(n), true
}

func (n Number) Bool() (bool, bool) {
	return n == 1, false
}

func (n Number) String() (string, bool) {
	return "", false
}

func (n Number) Object() (Object, bool) {
	return nil, false
}

func (n Number) Array() (Array, bool) {
	return nil, false
}

func (n Number) Type() Type {
	return NumberType
}

func (n Number) Clone() Value {
	return n
}

type Boolean bool

func (n Boolean) Int() (int, bool) {
	return 0, false
}

func (n Boolean) Float() (float64, bool) {
	return 0.0, false
}

func (n Boolean) Bool() (bool, bool) {
	return bool(n), false
}

func (n Boolean) String() (string, bool) {
	return "", false
}

func (n Boolean) Array() (Array, bool) {
	return nil, false
}

func (n Boolean) Object() (Object, bool) {
	return nil, false
}

func (n Boolean) Type() Type {
	return BooleanType
}

func (n Boolean) Clone() Value {
	return n
}

type String string

func (s String) Int() (int, bool) {
	return 0, false
}

func (s String) Float() (float64, bool) {
	return 0.0, false
}

func (s String) Bool() (bool, bool) {
	return false, false
}

func (s String) String() (string, bool) {
	return string(s), true
}

func (s String) Object() (Object, bool) {
	return nil, false
}

func (s String) Array() (Array, bool) {
	return nil, false
}

func (s String) Type() Type {
	return StringType
}

func (s String) Clone() Value {
	return s
}

type Array []Value

func (arr Array) Int() (int, bool) {
	return 0, false
}

func (arr Array) Float() (float64, bool) {
	return 0.0, false
}

func (arr Array) Bool() (bool, bool) {
	return false, false
}

func (arr Array) String() (string, bool) {
	return "", false
}

func (arr Array) Array() (Array, bool) {
	return arr, true
}

func (arr Array) Object() (Object, bool) {
	return nil, false
}

func (arr Array) Type() Type {
	return ArrayType
}

func (arr Array) Clone() Value {
	newArr := make(Array, arr.Len())
	for i, value := range arr {
		newArr[i] = value.Clone()
	}
	return newArr
}

func (arr Array) Children() []Value {
	return arr
}

func (arr Array) Len() int {
	return len(arr)
}

func (arr Array) GetChild(n int) (Value, bool) {
	if n < 0 || n >= arr.Len() {
		return nil, false
	}
	return arr[n], true
}

type Object map[string]Value

func (obj Object) Int() (int, bool) {
	return 0, false
}

func (obj Object) Float() (float64, bool) {
	return 0.0, false
}

func (obj Object) Bool() (bool, bool) {
	return false, false
}

func (obj Object) String() (string, bool) {
	return "", false
}

func (obj Object) Array() (Array, bool) {
	return nil, false
}

func (obj Object) Object() (Object, bool) {
	return obj, true
}

func (obj Object) Type() Type {
	return ObjectType
}

func (obj Object) Clone() Value {
	newMap := make(Object)
	for k, v := range obj {
		newMap[k] = v.Clone()
	}
	return newMap
}

func (obj Object) Get(pointer string) (Value, bool) {
	paths := strings.Split(pointer, ".")
	n := len(paths)
	objPointer := obj
	for i, path := range paths {
		if i == n-1 {
			v, ok := objPointer[path]
			return v, ok
		}

		if v, ok := objPointer[path]; !ok {
			return nil, false
		} else {
			if nextObj, ok := v.Object(); !ok {
				return nil, false
			} else {
				objPointer = nextObj
			}
		}
	}
	panic("unreachable")
}

func (obj Object) Map() map[string]Value {
	return obj
}
