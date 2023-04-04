package system

import (
	"fmt"
	"things/base/utils"
)

type ActorType int

const (
	PERSISTENCE ActorType = iota
	CONNECTIONS
	GATEWAY
	CORE
)

type System struct {
	Persistence Actor
	Connections Actor
	Gateway     Actor
	Core        Actor
}

var things System

func Register(t ActorType, actor Actor) {
	switch t {
	case PERSISTENCE:
		things.Persistence = actor
	case CONNECTIONS:
		things.Connections = actor
	case GATEWAY:
		things.Gateway = actor
	case CORE:
		things.Core = actor
	}
	go func() {
		if err := actor.Execute(); err != nil {
			utils.Logger.Fatalf("Ending actor %s: %s", actor.Name(), err.Error())
		}
	}()
}

func FindActor(t ActorType) (ret Actor, err error) {
	switch t {
	case PERSISTENCE:
		ret = things.Persistence
	case CONNECTIONS:
		ret = things.Connections
	case GATEWAY:
		ret = things.Gateway
	case CORE:
		ret = things.Core
	}
	if ret == nil {
		err = fmt.Errorf("actor type <%d> is not registered yet", t)
	}
	return
}

func MailboxOf(t ActorType) <-chan Message {
	actor, err := FindActor(t)
	if err != nil {
		return nil
	}
	return actor.Mailbox()
}

func Tell(t ActorType, msg any) error {
	actor, err := FindActor(t)
	if err != nil {
		return err
	}
	actor.Tell(msg)
	return nil
}

func Ask(t ActorType, msg any, future *utils.Future) error {
	actor, err := FindActor(t)
	if err != nil {
		return err
	}
	actor.Ask(msg, future)
	return nil
}
