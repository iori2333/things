package models

import (
	ejson "encoding/json"
	"fmt"
	"things/base/json"
	"things/base/utils"
)

// Features is a collection of thing Features.
// It can be designed as whether mutable or immutable.
type Features json.Object

func (fs Features) Create(name string, feature json.Value) (Features, error) {
	if _, ok := fs[name]; ok {
		return nil, fmt.Errorf("feature %s already exists", name)
	}
	ret := fs.Clone()
	ret[name] = feature
	return ret, nil
}

func (fs Features) Modify(name string, feature json.Value) (Features, json.Value) {
	ret := fs.Clone()
	feats, ok := ret[name]
	if !ok {
		ret[name] = feature
		return ret, feature
	}
	obj1, ok1 := feature.Object()
	obj2, ok2 := feats.Object()
	if ok1 && ok2 {
		json.MergeObject(obj1, obj2)
		ret[name] = obj1
	} else {
		ret[name] = feature
	}
	return ret, feats
}

func (fs Features) Overwrite(name string, feature json.Value) Features {
	ret := fs.Clone()
	ret[name] = feature
	return ret
}

func (fs Features) Delete(name string) Features {
	ret := fs.Clone()
	delete(ret, name)
	return ret
}

func (fs Features) Get(name string) (f json.Value, ok bool) {
	f, ok = fs[name]
	return
}

func (fs Features) Clone() Features {
	return utils.CloneMap(fs)
}

func (fs *Features) UnmarshalJSON(data []byte) error {
	obj := make(map[string]any)
	if err := ejson.Unmarshal(data, &obj); err != nil {
		return err
	}
	objs, err := json.ToObject(obj)
	if err != nil {
		return err
	}
	*fs = Features(objs)
	return nil
}
