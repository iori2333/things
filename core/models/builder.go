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
	b.features[name] = value
}

func (b *ThingBuilder) SetStates(states []State) {
	stateMap := make(States)
	for _, state := range states {
		stateMap[state.Name] = state.Transitions
	}
	b.states = stateMap
}

func (b *ThingBuilder) AppendState(name string, transitions []*Transition) {
	m := make(map[string]*Transition)
	for _, t := range transitions {
		m[t.Name] = t
	}
	if b.states == nil {
		b.states = make(States)
	}
	b.states[name] = m
}

func (b *ThingBuilder) SetInitialState(state string) {
	b.initialState = state
}

func (b *ThingBuilder) AppendTransition(state string, transition *Transition) {
	if b.states == nil {
		b.states = make(States)
	}
	if s, ok := b.states[state]; !ok {
		b.AppendState(state, []*Transition{transition})
	} else {
		s[transition.Name] = transition
	}
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
