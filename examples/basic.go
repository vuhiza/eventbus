package main

import (
	"fmt"
	"github.com/vuhiza/eventbus"
	"time"
)

type UserCreated struct {
	id   int
	name string
}

type UserDeleted struct {
	id   int
	name string
}

func OnUserCreated(event UserCreated) {
	fmt.Println("OnUserCreated", event)
}
func LogUserCreation(event UserCreated) {
	fmt.Println("User was created:", event)
}
func OnUserDeleted(event UserDeleted) {
	fmt.Println("OnUserDeleted", event)
}

func main() {

	bus := eventbus.NewEventBus()

	eventbus.Subscribe(bus, OnUserCreated)
	eventbus.Subscribe(bus, LogUserCreation)
	eventbus.Subscribe(bus, OnUserDeleted)

	bus.Publish(UserCreated{1, "John1"})
	bus.Publish(UserCreated{2, "John2"})
	bus.Publish(UserCreated{3, "John3"})
	bus.Publish(UserCreated{4, "John4"})
	bus.Publish(UserDeleted{1, "Doe1"})

	time.Sleep(10 * time.Second)
}
