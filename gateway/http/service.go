package http

import (
	"context"
	"fmt"

	"things/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Config struct {
	Port      int               `yaml:"port"`
	BasicAuth map[string]string `yaml:"basic_auth"`
}

// @Title                      Things HTTP API
// @Description                Access things, connections via HTTP.
// @SecurityDefinitions.Basic  BasicAuth
func Start(ctx context.Context, config *Config) error {
	engine := gin.Default()
	if config.Port == 0 {
		config.Port = 8080
	}

	docs.SwaggerInfo.BasePath = BasePath
	docs.SwaggerInfo.Version = Version

	group := engine.Group(BasePath)
	if config.BasicAuth != nil {
		group.Use(gin.BasicAuth(config.BasicAuth))
	}

	Register(group)
	engine.GET("/apidocs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, func(c *ginSwagger.Config) {
		c.Title = "Things HTTP API"
	}))

	if err := engine.Run(fmt.Sprintf(":%d", config.Port)); err != nil {
		return err
	}
	return nil
}
