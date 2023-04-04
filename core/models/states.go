package models

import (
	"fmt"
	"things/base/utils"
)

// State stores current state and its available transitions, which describes the behavior of a thing.
type State struct {
	Name        string                 `json:"name"`
	Transitions map[string]*Transition `json:"transitions"`
}

func (s *State) GetTransition(name string, input Message) (*Transition, error) {
	ret, ok := s.Transitions[name]
	if ok {
		return nil, fmt.Errorf("transition %s not found", name)
	}
	if err := ret.Validate(input); err != nil {
		return nil, err
	}
	return ret, nil
}

func (s *State) Clone() *State {
	return &State{Name: s.Name, Transitions: utils.CloneMap(s.Transitions)}
}

type States map[string]*State

func (ss States) Clone() States {
	return utils.CloneMap(ss)
}
