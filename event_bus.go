package eventbus

import (
	"context"
	"reflect"
	"sync"
)

type handlerFn = func(any)

type EventBus struct {
	mu        sync.RWMutex
	listeners map[reflect.Type][]handlerFn
	queues    map[reflect.Type]chan any
	ctx       context.Context
	cancel    context.CancelFunc
}

func NewEventBus() *EventBus {
	ctx, cancel := context.WithCancel(context.Background())
	return &EventBus{
		listeners: make(map[reflect.Type][]handlerFn),
		queues:    make(map[reflect.Type]chan any),
		ctx:       ctx,
		cancel:    cancel,
	}
}

func Subscribe[T any](bus *EventBus, handler func(T)) {
	eventType := reflect.TypeOf((*T)(nil)).Elem() // type of first parameter

	bus.mu.Lock()
	defer bus.mu.Unlock()

	if _, ok := bus.queues[eventType]; !ok {
		queue := make(chan any, 64)
		bus.queues[eventType] = queue
		go bus.fanOut(eventType, queue)
	}

	// Wrapper for generic typing
	bus.listeners[eventType] = append(bus.listeners[eventType], func(e any) {
		if eventValue, ok := e.(T); ok {
			handler(eventValue)
		}
	})
}

func (bus *EventBus) Publish(event any) {
	eventType := reflect.TypeOf(event)

	bus.mu.RLock()
	q, ok := bus.queues[eventType]
	bus.mu.RUnlock()
	if !ok {
		return
	}

	select {
	case q <- event:
	case <-bus.ctx.Done():
	}
}

func (bus *EventBus) Close() { bus.cancel() }

func (bus *EventBus) fanOut(t reflect.Type, queue <-chan any) {
	for {
		select {
		case event := <-queue:
			handlers := append([]handlerFn(nil), bus.listeners[t]...)

			for _, handler := range handlers {
				go handler(event)
			}
		case <-bus.ctx.Done():
			return
		}
	}
}
