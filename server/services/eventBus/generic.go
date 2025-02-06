package eventBus

type Handler[T any] func(data T)
type GlobalHandler[T any] func(event string, data T)

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

func (eb *EventBus[T]) Publish(event string, data T) {
	if handlers, found := eb.handlers[event]; found {
		for _, handler := range handlers {
			handler(data)
		}
	}
	for _, handler := range eb.globalHandlers {
		handler(event, data)
	}
}
