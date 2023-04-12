package v1

import (
	"things/base/system"
	"things/base/utils"
	"things/core/commands"
	"things/core/service"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	RetrievedThings = "Successfully retrieved things"
	RetrievedThing  = "Successfully retrieved thing details"
	CreatedThing    = "Successfully created thing"
	ReplacedThing   = "Successfully replaced thing"
	UpdatedThing    = "Successfully updated thing"
	DeletedThing    = "Successfully deleted thing"
	NotFound        = "Request entity not found"
	InvalidRequest  = "Invalid request parameters"
)

type ThingRoutes struct{}

// ListThings godoc
// @Summary List all things
// @Schemes
// @Description List all things and thing details stored in system
// @Tags Things
// @Produce json
// @Success 200 {array} Thing "Retrieved Things"
// @Router /things [get]
func (ThingRoutes) ListThings(c *gin.Context) {
	future := utils.NewFuture(10 * time.Second)

	system.Ask(system.CORE, &service.CommandMessage{
		Command: &commands.QueryThings{},
	}, future)

	res, err := future.Wait()
	if err != nil {
		c.AbortWithError(500, err)
	} else {
		c.JSON(200, res)
	}
}

// ListThings godoc
// @Summary List thing details
// @Schemes
// @Description List details of a thing stored in system
// @Tags Things
// @Param	namespace path string true "namespace of the thing"
// @Param	id path string true "id of the thing"
// @Produce json
// @Success 200 {object} Thing "Retrieved Thing"
// @Failure 400 {object} Error "Invalid Request"
// @Failure 404 {object} Error "NotFound"
// @Router /things/{namespace}/{id} [get]
func (ThingRoutes) ListThing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// CreateThing godoc
// @Summary Create a thing
// @Schemes
// @Description Create a new thing
// @Param	thing body Thing true "thing to create"
// @Tags Things
// @Accept json
// @Produce json
// @Success 201 {object} Thing "Created Thing"
// @Failure 400 {object} Error "Invalid Request"
// @Router /things [post]
func (ThingRoutes) CreateThing(c *gin.Context) {
	// TODO: create thing
}

// ReplaceThing godoc
// @Summary Replace thing properties
// @Schemes
// @Description Replace thing properties. ThingID and Namespace cannot be changed, thus they are ignored. If thing to replace does not exist, it will be created.
// @Tags Things
// @Param	namespace path string true "namespace of the thing"
// @Param	id path string true "id of the thing"
// @Param	thing body Thing true "new properties of the thing"
// @Accept json
// @Produce json
// @Success 200 {object} Thing "Replaced Thing"
// @Success 201 {object} Thing "Created Thing"
// @Failure 400 {object} Error "Invalid Request"
// @Router /things/{namespace}/{id} [post]
func (ThingRoutes) ReplaceThing(c *gin.Context) {
	// TODO: replace thing
}

// UpdateThing godoc
// @Summary Update thing properties
// @Schemes
// @Description Update thing properties. ThingID and Namespace cannot be changed, thus they are ignored. If thing to update does not exist, it will be created.
// @Tags Things
// @Param	namespace path string true "namespace of the thing"
// @Param	id path string true "id of the thing"
// @Param	thing body Thing true "updated properties of the thing"
// @Accept json
// @Produce json
// @Success 200 {object} Thing "Updated Thing"
// @Success 201 {object} Thing "Created Thing"
// @Failure 400 {object} Error "Invalid Request"
// @Router /things/{namespace}/{id} [put]
func (ThingRoutes) UpdateThing(c *gin.Context) {
	// TODO: update thing
}

// DeleteThing godoc
// @Summary Delete a thing
// @Schemes
// @Description Delete a thing from the system
// @Tags Things
// @Param	namespace path string true "namespace of the thing"
// @Param	id path string true "id of the thing"
// @Accept json
// @Produce json
// @Success 200 {object} Thing "Updated Thing"
// @Success 201 {object} Thing "Created Thing"
// @Failure 400 {object} Error "Invalid Request"
// @Router /things/{namespace}/{id} [delete]
func (ThingRoutes) DeleteThing(c *gin.Context) {
	// TODO: delete thing
}

var thing ThingRoutes
