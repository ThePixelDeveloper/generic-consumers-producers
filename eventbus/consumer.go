package eventbus

import "context"

type Consumer[T Message] struct {
	Queue    Memory
	handlers map[string]HandlerFunc[T]
}

type HandlerFunc[T Message] func(ctx context.Context, event T) error

func (c *Consumer[T]) On(e T, h HandlerFunc[T]) error {
	if c.handlers == nil {
		c.handlers = make(map[string]HandlerFunc[T])
	}

	c.handlers[e.Topic()] = h
	return nil
}

func (c *Consumer[T]) Listen() {
	go func() {
		for {
			select {
			case msg := <-c.Queue.Receive(context.Background()):
				if handler, ok := c.handlers[msg.Topic()]; ok {
					handler(context.Background(), msg.(T))
				}
			}
		}
	}()
}
