package eventbus

import (
	"context"
)

type Message interface {
	Topic() string
}

type Memory struct {
	Messages chan Message
}

func (m Memory) Send(ctx context.Context, message Message) {
	m.Messages <- message
}

func (m Memory) Receive(ctx context.Context) <-chan Message {
	return m.Messages
}
