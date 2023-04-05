package models

import (
	"fmt"
	"things/base/json"
	"things/base/utils"
)

type ThingId struct {
	Namespace string `json:"namespace"`
	Id        string `json:"id"`
}

func (id *ThingId) String() string {
	return id.Namespace + ":" + id.Id
}

type Message = json.Object

type MessageType map[string]json.Type

func (mt MessageType) Clone() MessageType {
	return utils.CloneMap(mt)
}

type Thing struct {
	ThingId   ThingId  `json:"thing_id"`
	Features  Features `json:"features,omitempty"`
	States    States   `json:"states"`
	StateName string   `json:"state"`
}

func (t *Thing) Namespace() string {
	return t.ThingId.Namespace
}

func (t *Thing) State() Transitions {
	return t.States[t.StateName]
}

func (t *Thing) Tell(name string, msg Message) error {
	state := t.State()
	transition, err := state.GetTransition(name, msg)
	if err != nil {
		return err
	}
	t.StateName = transition.NextState
	return nil
}

func (t *Thing) Ask(name string, msg Message, future *utils.Future) {
	state := t.State()
	transition, err := state.GetTransition(name, msg)
	if err != nil {
		future.SetError(err)
	}
	future.OnResult(func(v any) { t.StateName = transition.NextState })
}

func (t *Thing) Clone() *Thing {
	return &Thing{
		ThingId:   t.ThingId,
		Features:  t.Features.Clone(),
		StateName: t.StateName,
		States:    t.States.Clone(),
	}
}

func (t *Thing) Validate() error {
	if t.ThingId.Namespace == "" || t.ThingId.Id == "" {
		return fmt.Errorf("invalid thing id: %v", t.ThingId)
	}

	if _, ok := t.States[t.StateName]; !ok {
		return fmt.Errorf("initial state %s not found", t.StateName)
	}

	for _, state := range t.States {
		for _, transition := range state {
			if _, ok := t.States[transition.NextState]; !ok {
				return fmt.Errorf("state %s not found", transition.NextState)
			}
		}
	}
	return nil
}
