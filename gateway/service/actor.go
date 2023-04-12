package service

import (
	"context"
	"things/base/system"
	"things/gateway/http"
)

type Config struct {
	Enabled bool        `yaml:"enabled"`
	HTTP    http.Config `yaml:"http"`
}

type Actor struct {
	system.BaseActor[Config]
}

func (a *Actor) Execute() error {
	return nil
}

func Start(ctx context.Context, config *Config) {
	if !config.Enabled {
		return
	}
	actor := &Actor{
		BaseActor: system.MakeActor("GATEWAY", ctx, config),
	}
	system.Register(system.GATEWAY, actor)
	go func() {
		if err := http.Start(ctx, &config.HTTP); err != nil {
			actor.Logger.Printf("HTTP service failed unexpectedly: %s", err.Error())
		}
	}()
	actor.Logger.Println("Started Gateway service")
}
