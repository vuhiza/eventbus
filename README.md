# EventBus

A lightweight, type-safe, generic event bus for Go.

This library provides a simple publish/subscribe mechanism using Go generics and reflection. Handlers are called
concurrently in their own goroutines.

## ✨ Features

- Type-safe subscriptions using generics
- Automatic fan-out to multiple subscribers
- Buffered event queues
- Context-based cancellation
- Simple API

## 🚀 Usage

```go
package main

import (
	"fmt"
	"time"

	"github.com/vuhiza/eventbus"
)

type MyEvent struct {
	Message string
}

func main() {
	bus := eventbus.NewEventBus()

	// Subscribe to MyEvent
	eventbus.Subscribe(bus, func(e MyEvent) {
		fmt.Println("Received MyEvent:", e.Message)
	})

	// Publish an event
	bus.Publish(MyEvent{Message: "Hello EventBus!"})

	// Give goroutines time to process
	time.Sleep(100 * time.Millisecond)

	// Close the bus when done
	bus.Close()
}
```