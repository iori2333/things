package v1

import (
	"github.com/gin-gonic/gin"
)

const APIVersion = "/api/v1"

type Route interface {
	Register(group *gin.RouterGroup)
}

type SubRoute map[string]Route

func (r SubRoute) Register(group *gin.RouterGroup) {
	for path, route := range r {
		route.Register(group.Group(path))
	}
}

type EndPointRoute map[string]gin.HandlerFunc

func (r EndPointRoute) Register(group *gin.RouterGroup) {
	for method, handler := range r {
		group.Handle(method, "/", handler)
	}
}

type S = SubRoute
type E = EndPointRoute

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type API string

const (
	Things      API = "/things"
	Connections API = "/connections"
)

var Routes = map[API]Route{
	Things: S{
		"/": E{
			GET:  thing.ListThings,
			POST: thing.CreateThing,
		},
		"/:ns/:id": S{
			"/": E{
				GET:    thing.ListThing,
				POST:   thing.ReplaceThing,
				PUT:    thing.UpdateThing,
				DELETE: thing.DeleteThing,
			},
			"/features":     S{},
			"/states":       E{},
			"/state/:name":  S{},
			"/message/:msg": E{},
		},
	},
	Connections: E{},
}

func Register(group *gin.RouterGroup) {
	for path, route := range Routes {
		route.Register(group.Group(string(path)))
	}
}
