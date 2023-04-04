package service

import (
	"errors"
	"things/base/system"
	"things/base/utils"
	"things/core/commands"
)

type Message struct {
	Type commands.Type
	Data []byte
}

func (msg *Message) Command() (commands.Interface, error) {
	return commands.Unmarshal(msg.Type, msg.Data)
}

type CommandHandler struct{}

func (handler CommandHandler) convert(msg system.Message) (*Message, *utils.Future, error) {
	v, ok := msg.Payload.(Message)
	if !ok || msg.Future == nil {
		return nil, nil, errors.New("ignoring invalid message")
	}
	return &v, msg.Future, nil
}

func (handler CommandHandler) Execute(sysMsg system.Message) {
	msg, future, err := handler.convert(sysMsg)
	if err != nil {
		return
	}

	if cmd, err := msg.Command(); err != nil {
		future.SetError(err)
	} else {
		cmd.Execute(future)
	}
}
