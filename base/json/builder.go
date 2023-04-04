package json

import "strings"

func ToObject(obj map[string]any) (Object, error) {
	ret := make(Object)
	for k, v := range obj {
		value, err := ToValue(v)
		if err != nil {
			return nil, err
		}
		ret[k] = value
	}
	return ret, nil
}

func ToArray(arr []any) (Array, error) {
	ret := make(Array, len(arr))
	for i, v := range arr {
		value, err := ToValue(v)
		if err != nil {
			return nil, err
		}
		ret[i] = value
	}
	return ret, nil
}

func ToValue(v any) (Value, error) {
	if v == nil {
		return nil, nil
	}
	switch v := v.(type) {
	case int:
		return Number(v), nil
	case float64:
		return Number(v), nil
	case bool:
		return Boolean(v), nil
	case string:
		return String(v), nil
	case map[string]any:
		return ToObject(v)
	case []any:
		return ToArray(v)
	case Value:
		return v, nil
	default:
		return nil, nil
	}
}

func MergeObject(obj1, obj2 Object) {
	for k, v := range obj2 {
		if value, ok := obj1[k]; !ok {
			obj1[k] = v
		} else {
			nextObj, ok1 := value.Object()
			thisObj, ok2 := v.Object()
			if !ok1 || !ok2 {
				obj1[k] = v
			} else {
				MergeObject(nextObj, thisObj)
			}
		}
	}
}

// ValueBuilder builds a single json value from existing types
type ValueBuilder struct {
	value Value
}

func (f *ValueBuilder) Set(v any) error {
	value, err := ToValue(v)
	if err != nil {
		return err
	}
	f.value = value
	return nil
}

func (f *ValueBuilder) SetInt(n int) {
	f.value = Number(n)
}

func (f *ValueBuilder) SetFloat(n float64) {
	f.value = Number(n)
}

func (f *ValueBuilder) SetBoolean(n bool) {
	f.value = Boolean(n)
}

func (f *ValueBuilder) SetString(s string) {
	f.value = String(s)
}

// SetValue consumes Value references instead of their clones.
// You may use Value.Clone() to prevent modifications
func (f *ValueBuilder) SetValue(v Value) {
	f.value = v
}

func (f *ValueBuilder) Build() Value {
	return f.value
}

type ArrayBuilder struct {
	children []Value
}

func (f *ArrayBuilder) AppendValue(v Value) {
	f.children = append(f.children, v)
}

func (f *ArrayBuilder) Append(v any) error {
	value, err := ToValue(v)
	if err != nil {
		return err
	}
	f.AppendValue(value)
	return nil
}

func (f *ArrayBuilder) BuildArray() Array {
	return f.children
}

func (f *ArrayBuilder) Build() Value {
	return f.BuildArray()
}

type ObjectBuilder struct {
	object map[string]Value
}

func (f *ObjectBuilder) SetObject(object Object) {
	f.object = object
}

func (f *ObjectBuilder) Merge(object Object) {
	MergeObject(f.object, object)
}

func (f *ObjectBuilder) SetValue(path string, value Value) {
	if f.object == nil {
		f.object = make(map[string]Value)
	}
	paths := strings.Split(path, ".")
	n := len(paths)
	pointer := f.object
	for i, path := range paths {
		if i == n-1 {
			pointer[path] = value
			return
		}

		v, ok := pointer[path]
		var newObj Object
		if !ok || v == nil {
			// meet null or undefined property
			newObj = Object{}
			pointer[path] = newObj
		} else if newObj, ok = v.Object(); !ok {
			// if property exists and not an object, overwrite it
			newObj = Object{}
			pointer[path] = newObj
		}
		pointer = newObj
	}
}

func (f *ObjectBuilder) Set(path string, v any) error {
	value, err := ToValue(v)
	if err != nil {
		return err
	}
	f.SetValue(path, value)
	return nil
}

func (f *ObjectBuilder) BuildObject() Object {
	return f.object
}

func (f *ObjectBuilder) Build() Value {
	return f.BuildObject()
}
