package main

import (
	"fmt"
	"github.com/vuhiza/eventbus"
	"time"
)

type ExampleEvent struct {
	id      int
	message string
}

type EventHandlerContainer struct {
	dependency string
}

func (d EventHandlerContainer) Event1Handler(event ExampleEvent) {
	fmt.Println("Example event was handled", event, d.dependency)

}

func main() {

	bus := eventbus.NewEventBus()

	eventHandler := EventHandlerContainer{
		dependency: "Important dependency",
	}

	eventbus.Subscribe(bus, eventHandler.Event1Handler)

	bus.Publish(ExampleEvent{id: 3, message: "hello world"})

	time.Sleep(1 * time.Second)

	bus.Close()

}
