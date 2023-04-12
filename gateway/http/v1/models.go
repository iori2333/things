package v1

import (
	"things/base/errors"
	things "things/core/models"
)

type Thing struct {
	things.Thing
	States map[string][]things.Transition `json:"states"`
}

type Error = errors.Error
