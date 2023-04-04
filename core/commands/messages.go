package commands

import (
	"things/base/json"
	"things/base/utils"
	"things/core/models"
)

type TellThing struct {
	ThingId models.ThingId `json:"thing_id"`
	Message json.Object    `json:"message"`
}

func (c *TellThing) CommandType() Type {
	return TypeTellThing
}

func (c *TellThing) Execute(future *utils.Future) {
}

type AskThing struct {
	ThingId models.ThingId `json:"thing_id"`
	Message json.Object    `json:"message"`
}

func (c *AskThing) CommandType() Type {
	return TypeAskThing
}

func (c *AskThing) Execute(future *utils.Future) {
}
