package service

import (
	"context"
	"fmt"
	"sync"
	"things/base/system"
	db "things/persistence/database"
)

type Config struct {
	Database map[string]any `yaml:"database"`
}

type Actor struct {
	system.BaseActor[Config]
	Database db.Interface
}

func Start(ctx context.Context, config *Config) {
	actor := &Actor{
		BaseActor: system.MakeActor("PERSISTENCE", ctx, config),
	}
	actor.Database = actor.GetDatabase(config.Database)
	system.Register(system.PERSISTENCE, actor)
	actor.Logger.Println("Started Persistence service")
	<-ctx.Done()
}

func (a *Actor) Execute() error {
	for {
		select {
		case <-a.Context.Done():
			return nil
		case _, ok := <-a.Mailbox():
			if !ok {
				return fmt.Errorf("sending message to dead actor: %s", a.Name())
			}
		}
	}
}

// db interfaces
var (
	once     sync.Once
	instance db.Interface
)
