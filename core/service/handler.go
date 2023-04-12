package service

import (
	"things/base/errors"
	"things/base/system"
	"things/base/utils"
	"things/core/commands"
)

type Message interface {
	GetCommand() (commands.Interface, error)
}

type BytesMessage struct {
	Type commands.Type
	Data []byte
}

func (msg BytesMessage) GetCommand() (commands.Interface, error) {
	return commands.Unmarshal(msg.Type, msg.Data)
}

type CommandMessage struct {
	Command commands.Interface
}

func (msg CommandMessage) GetCommand() (commands.Interface, error) {
	return msg.Command, nil
}

type CommandHandler struct{}

func (handler CommandHandler) convert(msg system.Message) (Message, *utils.Future, error) {
	v, ok := msg.Payload.(Message)
	if !ok || msg.Future == nil {
		return nil, msg.Future, errors.Default("InvalidMessage", "Future is nil or payload is not a Message")
	}
	return v, msg.Future, nil
}

func (handler CommandHandler) Execute(sysMsg system.Message) {
	msg, future, err := handler.convert(sysMsg)
	if err != nil {
		if future != nil {
			future.SetError(err)
		}
		return
	}

	if cmd, err := msg.GetCommand(); err != nil {
		future.SetError(err)
	} else {
		cmd.Execute(future)
	}
}
