package models

import "things/base/json"

type ThingBuilder struct {
	thingId      ThingId
	features     Features
	states       States
	initialState string
}

func (b *ThingBuilder) SetThingId(id ThingId) {
	b.thingId = id
}

func (b *ThingBuilder) SetFeatures(features Features) {
	b.features = features
}

func (b *ThingBuilder) AppendFeature(name string, value json.Object) {
	if b.features == nil {
		b.features = make(Features)
	}
	b.features[name] = &Feature{
		Id:         name,
		Properties: value,
	}
}

func (b *ThingBuilder) SetStates(states States) {
	b.states = states
}

func (b *ThingBuilder) AppendState(name string, transitions []*Transition) {
	m := make(map[string]*Transition)
	for _, t := range transitions {
		m[t.Name] = t
	}
	if b.states == nil {
		b.states = make(States)
	}
	b.states[name] = &State{
		Name:        name,
		Transitions: m,
	}
}

func (b *ThingBuilder) SetInitialState(state string) {
	b.initialState = state
}

func (b *ThingBuilder) Build() (*Thing, error) {
	thing := &Thing{
		ThingId:   b.thingId,
		Features:  b.features,
		States:    b.states,
		StateName: b.initialState,
	}
	if err := thing.Validate(); err != nil {
		return nil, err
	}
	return thing, nil
}
