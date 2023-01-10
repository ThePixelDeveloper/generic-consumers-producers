package eventbus

import (
	"context"
)

type Producer struct {
	Queue Memory
}

func (e *Producer) Emit(ctx context.Context, event Message) {
	go e.Queue.Send(ctx, event)
}
