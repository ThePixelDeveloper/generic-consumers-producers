package eventbus

import (
	"context"
)

type Producer[T Message] struct {
	Queue Memory
}

func (e *Producer[T]) Emit(ctx context.Context, event T) error {
	go e.Queue.Send(ctx, event)
	return nil
}
