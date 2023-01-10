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
	var (
		err error
		wg  sync.WaitGroup
	)

	// First we'll setup an in memory queue; It's using channels, but it could also be SQS, Kafka, etc.
	queue := eventbus.Memory{
		Messages: make(chan eventbus.Message),
	}

	// Potentially an idea on how we'd define a generic producer
	userCreatedProducer := eventbus.Producer[events.UserCreated]{Queue: queue}

	// Potentially an idea on how we'd define a generic consumer
	userCreatedConsumer := eventbus.Consumer[events.UserCreated]{Queue: queue}
	userCreatedConsumer.Listen()

	wg.Add(1)
	err = userCreatedProducer.Emit(context.Background(), events.UserCreated{
		ID:   1,
		Name: "John Doe",
	})
	if err != nil {
		panic(err)
	}

	err = userCreatedConsumer.On(events.UserCreated{}, func(ctx context.Context, event events.UserCreated) error {
		fmt.Printf("User created: %+v", event)
		wg.Done()
		return nil
	})
	if err != nil {
		panic(err)
	}

	wg.Wait()
}
