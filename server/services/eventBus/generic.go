package eventBus

import "server/models"

type Handler[T any] func(data T, user models.User)
type GlobalHandler[T any] func(event string, data T, user models.User)

type EventBus[T any] struct {
	handlers       map[string][]Handler[T]
	globalHandlers []GlobalHandler[T]
}

func NewEventBus[T any]() *EventBus[T] {
	return &EventBus[T]{
		handlers:       make(map[string][]Handler[T]),
		globalHandlers: []GlobalHandler[T]{},
	}
}

func (eb *EventBus[T]) Subscribe(event string, handler Handler[T]) {
	eb.handlers[event] = append(eb.handlers[event], handler)
}

func (eb *EventBus[T]) SubscribeGlobal(handler GlobalHandler[T]) {
	eb.globalHandlers = append(eb.globalHandlers, handler)
}

func (eb *EventBus[T]) Publish(event string, data T, user models.User) {
	if handlers, found := eb.handlers[event]; found {
		for _, handler := range handlers {
			handler(data, user)
		}
	}
	for _, handler := range eb.globalHandlers {
		handler(event, data, user)
	}
}
