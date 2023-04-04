package service

import (
	"context"
	"fmt"
	"things/base/system"
)

type Config struct{}

type Actor struct {
	system.BaseActor[Config]
	Handler CommandHandler
}

func (a *Actor) Execute() error {
	for {
		select {
		case <-a.Context.Done():
			return nil
		case msg, ok := <-a.Mailbox():
			if !ok {
				return fmt.Errorf("sending message to dead actor: %s", a.Name())
			}
			go a.Handler.Execute(msg)
		}
	}
}

func Start(ctx context.Context, config *Config) {
	actor := &Actor{
		BaseActor: system.MakeActor("CORE", ctx, config),
		Handler:   CommandHandler{},
	}
	system.Register(system.CORE, actor)
	actor.Logger.Println("Started Core service")
	<-ctx.Done()
}
