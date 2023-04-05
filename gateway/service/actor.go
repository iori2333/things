package service

import (
	"context"
	"things/base/system"
)

type Config struct{}

type Actor struct {
	system.BaseActor[Config]
}

func (a *Actor) Execute() error {
	return nil
}

func Start(ctx context.Context, config *Config) {
	actor := &Actor{
		BaseActor: system.MakeActor("GATEWAY", ctx, config),
	}
	system.Register(system.CORE, actor)
	actor.Logger.Println("Started Gateway service")
}
