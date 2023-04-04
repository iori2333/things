package json

import (
	"encoding/json"
	"testing"
)

func TestJsonStringify(t *testing.T) {
	builder := ObjectBuilder{}
	_ = builder.Set("name", "name")
	_ = builder.Set("age", 20)
	obj1 := builder.BuildObject()

	builder2 := ObjectBuilder{}
	_ = builder2.Set("name2", "name")
	_ = builder2.Set("age", 28)
	obj2 := builder2.BuildObject()

	builder3 := ObjectBuilder{}
	builder3.SetObject(obj1)
	builder3.Merge(obj2)
	_ = builder3.BuildObject()

}

func TestJsonBuilder(t *testing.T) {
	obj := map[string]any{
		"name": "Iori",
		"age":  20,
		"info": nil,
	}
	jObj, err := ToValue(obj)
	if err != nil {
		panic(err)
	}
	o, err := json.Marshal(jObj)
	if err != nil {
		panic(err)
	}
	println(string(o))
}
