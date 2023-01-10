package eventbus

import "context"

type Consumer[T Message] struct {
	Queue   Memory
	handler HandlerFunc[T]
}

type HandlerFunc[T Message] func(ctx context.Context, event T)

func (c *Consumer[T]) Handle(h HandlerFunc[T]) {
	c.handler = h
}

func (c *Consumer[T]) Listen() {
	go func() {
		for {
			select {
			case msg := <-c.Queue.Receive(context.Background()):
				if _, ok := msg.(T); ok {
					c.handler(context.Background(), msg.(T))
				}
			}
		}
	}()
}
