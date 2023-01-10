package main

import (
	"context"
	"fmt"
	"generic-consumers-producers/eventbus"
	"generic-consumers-producers/events"
	"sync"
)

// A trying example of generic consumers and producers
func main() {
	var wg sync.WaitGroup

	// First we'll setup an in memory queue; It's using channels, but it could also be SQS, Kafka, etc.
	queue := eventbus.Memory{Messages: make(chan eventbus.Message)}

	wg.Add(1)

	// Producers can stay the same and don't need to be generic
	producer := eventbus.Producer{Queue: queue}
	producer.Emit(context.Background(), events.UserCreated{
		ID:   1,
		Name: "John Doe",
	})

	// Potentially an idea on how we'd define a generic consumer
	userCreatedConsumer := eventbus.Consumer[events.UserCreated]{Queue: queue}

	// The benefit over the current implementation is that we can remove the common boilerplate of
	// event.(events.UserCreated) and just use the generic type
	userCreatedConsumer.On(events.UserCreated{}, func(ctx context.Context, event events.UserCreated) {
		fmt.Printf("User created: %+v", event)
		wg.Done()
	})
	userCreatedConsumer.Listen()

	wg.Wait()
}
