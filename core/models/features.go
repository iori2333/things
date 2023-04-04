package models

import (
	"fmt"
	"things/base/json"
)

// Feature indicates any properties of things which can be changed over time.
// User can modify the value of a feature via commands.
// A feature is represented as a json object.
type Feature struct {
	Id         string
	Properties json.Object `json:"properties"`
}

func (f *Feature) Select(path string) (json.Value, bool) {
	obj, ok := f.Properties.Object()
	if !ok {
		return nil, false
	}
	return obj.Get(path)
}

func (f *Feature) Clone() *Feature {
	return &Feature{
		Id:         f.Id,
		Properties: f.Properties.Clone().(json.Object),
	}
}

// Features is a collection of thing Features.
// It can be designed as whether mutable or immutable.
type Features map[string]*Feature

func (fs Features) Create(name string, feature *Feature) (Features, error) {
	if _, ok := fs[name]; ok {
		return nil, fmt.Errorf("feature %s already exists", name)
	}
	ret := fs.Clone()
	ret[name] = feature
	return ret, nil
}

func (fs Features) Modify(name string, feature *Feature) (Features, *Feature) {
	ret := fs.Clone()
	if feats, ok := fs[name]; !ok {
		ret[name] = feature
		return ret, feature
	} else {
		json.MergeObject(feats.Properties, feature.Properties)
		return ret, feats
	}
}

func (fs Features) Overwrite(name string, feature *Feature) Features {
	ret := fs.Clone()
	ret[name] = feature
	return ret
}

func (fs Features) Delete(name string) Features {
	ret := fs.Clone()
	delete(ret, name)
	return ret
}

func (fs Features) Get(name string) (f *Feature, ok bool) {
	f, ok = fs[name]
	return
}

func (fs Features) List() (ret []*Feature) {
	for _, v := range fs {
		ret = append(ret, v)
	}
	return
}

func (fs Features) Clone() Features {
	ret := make(Features)
	for k, v := range fs {
		ret[k] = v.Clone()
	}
	return ret
}
