package models

import (
	"encoding/json"
	"fmt"
	"things/base/utils"
)

// Transition defines the behavior of a thing
// let f <- Transition, x <- Inputs, then NextState <- f(x)
type Transition struct {
	Name        string      `json:"name"`
	MessageType MessageType `json:"message_type"`
	NextState   string      `json:"next_state"`
}

func (t *Transition) Validate(input Message) error {
	for k, v := range t.MessageType {
		t, ok := input[k]
		if !ok {
			if v.IsNullable() {
				continue
			}
			return fmt.Errorf("missing input: %s", k)
		}
		if v != t.Type() {
			return fmt.Errorf("invalid input type for %s: %s", k, t.Type())
		}
	}
	return nil
}

func (t *Transition) Clone() *Transition {
	return &Transition{
		Name:        t.Name,
		MessageType: t.MessageType.Clone(),
		NextState:   t.NextState,
	}
}

// State is a serializable representation of a thing's state
type State struct {
	Name        string      `json:"name"`
	Transitions Transitions `json:"transitions"`
}

func (s *State) Clone() *State {
	return &State{Name: s.Name, Transitions: utils.CloneMap(s.Transitions)}
}

type Transitions map[string]*Transition

func (ts Transitions) Clone() Transitions {
	return utils.CloneMap(ts)
}

func (ts Transitions) GetTransition(name string, input Message) (*Transition, error) {
	ret, ok := ts[name]
	if ok {
		return nil, fmt.Errorf("transition %s not found", name)
	}
	if err := ret.Validate(input); err != nil {
		return nil, err
	}
	return ret, nil
}

func (ts *Transitions) UnmarshalJSON(data []byte) error {
	if *ts == nil {
		*ts = make(map[string]*Transition)
	}
	var transitions []*Transition
	if err := json.Unmarshal(data, &transitions); err != nil {
		return err
	}
	for _, t := range transitions {
		(*ts)[t.Name] = t
	}
	return nil
}

func (ts Transitions) MarshalJSON() ([]byte, error) {
	var transitions []*Transition
	for _, t := range ts {
		transitions = append(transitions, t)
	}
	return json.Marshal(transitions)
}

type States map[string]Transitions

func (ss States) Clone() States {
	return utils.CloneMap(ss)
}
