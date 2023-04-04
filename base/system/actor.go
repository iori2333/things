package system

import (
	"context"
	"log"
	"os"
	"things/base/utils"
)

type Actor interface {
	Name() string
	Mailbox() <-chan Message
	Tell(msg any)
	Ask(msg any, future *utils.Future)
	Execute() error
}

type Message struct {
	Payload any
	Future  *utils.Future
}

type MessageHandler interface {
	Execute(msg Message) bool
}

type BaseActor[T any] struct {
	mailbox chan Message
	name    string
	Logger  *log.Logger
	Context context.Context
	Config  *T
}

func MakeActor[T any](name string, ctx context.Context, config *T) BaseActor[T] {
	return BaseActor[T]{
		mailbox: make(chan Message),
		name:    name,
		Logger:  utils.GetLogger(name, os.Stderr),
		Context: ctx,
		Config:  config,
	}
}

func (a *BaseActor[T]) Name() string {
	return a.name
}

func (a *BaseActor[T]) Mailbox() <-chan Message {
	return a.mailbox
}

func (a *BaseActor[T]) Tell(msg any) {
	a.mailbox <- Message{Payload: msg}
}

func (a *BaseActor[T]) Ask(msg any, future *utils.Future) {
	a.mailbox <- Message{Payload: msg, Future: future}
}
