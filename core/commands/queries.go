package commands

import (
	"things/base/utils"
	"things/core/models"
)

type QueryThing models.ThingId

func (c *QueryThing) CommandType() Type {
	return TypeQueryThing
}

func (c *QueryThing) Execute(future *utils.Future) {

}

type QueryThings struct {
	Feature *models.Feature `json:"feature"`
	State   string          `json:"state"`
}

func (c *QueryThings) CommandType() Type {
	return TypeQueryThings
}

func (c *QueryThings) Execute(future *utils.Future) {

}

type QueryFeature struct {
	ThingId   models.Thing `json:"thing_id"`
	FeatureId string       `json:"feature_id"`
}

func (c *QueryFeature) CommandType() Type {
	return TypeQueryFeature
}

func (c *QueryFeature) Execute(future *utils.Future) {

}

type QueryFeatures struct {
	ThingId models.ThingId  `json:"thing_id"`
	Filter  *models.Feature `json:"filter"`
}

func (c *QueryFeatures) CommandType() Type {
	return TypeQueryFeatures
}

func (c *QueryFeatures) Execute(future *utils.Future) {

}

type QueryState models.ThingId

func (c *QueryState) CommandType() Type {
	return TypeQueryState
}

func (c *QueryState) Execute(future *utils.Future) {

}

type QueryStates models.ThingId

func (c *QueryStates) CommandType() Type {
	return TypeQueryStates
}

func (c *QueryStates) Execute(future *utils.Future) {

}
