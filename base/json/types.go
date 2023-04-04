package json

import "strings"

type Type string

const (
	NumberType  Type = "number"
	BooleanType Type = "boolean"
	StringType  Type = "string"
	ArrayType   Type = "Array"
	ObjectType  Type = "Object"
	Null             = "null"
)

func (t Type) IsNullable() bool {
	return strings.HasSuffix(t.String(), Null)
}

func (t Type) ToNullable() Type {
	if t.IsNullable() {
		return t
	}
	return t + " | null"
}

func (t Type) String() string {
	return string(t)
}

func (t Type) Clone() Type {
	return t
}
