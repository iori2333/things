package commands

import (
	"things/base/json"
	"things/base/system"
	"things/base/utils"
	"things/core/models"
)

type CreateThing models.Thing

func (c *CreateThing) CommandType() Type {
	return TypeCreateThing
}

func (c *CreateThing) Execute(future *utils.Future) {
	thing := &models.Thing{
		ThingId:   c.ThingId,
		Features:  c.Features,
		States:    c.States,
		StateName: c.StateName,
	}
	if err := thing.Validate(); err != nil {
		future.SetError(err)
	} else if err := system.Tell(system.PERSISTENCE, thing); err != nil {
		future.SetError(err)
	} else {
		future.SetResult(thing)
	}
}

type ModifyThing models.Thing

func (m *ModifyThing) CommandType() Type {
	return TypeModifyThing
}

func (m *ModifyThing) Execute(future *utils.Future) {
	// TODO: modify thing
}

type DeleteThing models.ThingId

func (d *DeleteThing) CommandType() Type {
	return TypeDeleteThing
}

func (d *DeleteThing) Execute(future *utils.Future) {
}

type CreateFeature struct {
	ThingId   models.ThingId `json:"thing_id"`
	FeatureId string         `json:"feature_id"`
	Feature   json.Value     `json:"feature"`
}

func (c *CreateFeature) CommandType() Type {
	return TypeCreateFeature
}

func (c *CreateFeature) Execute(future *utils.Future) {
	// TODO: create feature
}

type ModifyFeature struct {
	ThingId   models.ThingId `json:"thing_id"`
	FeatureId string         `json:"feature_id"`
	Path      string         `json:"path"`
	Value     json.Value     `json:"value"`
}

func (m *ModifyFeature) CommandType() Type {
	return TypeModifyFeature
}

func (m *ModifyFeature) Execute(future *utils.Future) {
	// TODO: modify feature

}

type OverwriteFeature struct {
	ThingId   models.ThingId `json:"thing_id"`
	FeatureId string         `json:"feature_id"`
	Feature   json.Value     `json:"feature"`
}

func (o *OverwriteFeature) CommandType() Type {
	return TypeOverwriteFeature
}

func (o *OverwriteFeature) Execute(future *utils.Future) {
	// TODO: overwrite feature
}

type DeleteFeature struct {
	ThingId   models.ThingId `json:"thing_id"`
	FeatureId string         `json:"feature_id"`
}

func (d *DeleteFeature) CommandType() Type {
	return TypeDeleteFeature
}

func (d *DeleteFeature) Execute(future *utils.Future) {
	// TODO: delete feature

}

type ModifyStates struct {
	ThingId models.ThingId `json:"thing_id"`
	States  models.States  `json:"states"`
}

func (m *ModifyStates) CommandType() Type {
	return TypeModifyStates
}

func (m *ModifyStates) Execute(future *utils.Future) {
	// TODO: modify states
}

type ModifyTransitions struct {
	ThingId     models.ThingId       `json:"thing_id"`
	StateName   string               `json:"state_name"`
	Transitions []*models.Transition `json:"transitions"`
}

func (m *ModifyTransitions) CommandType() Type {
	return TypeModifyTransitions
}

func (m *ModifyTransitions) Execute(future *utils.Future) {
	// TODO: modify transitions
}
