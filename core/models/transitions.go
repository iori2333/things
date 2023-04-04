package models

import (
	"fmt"
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
