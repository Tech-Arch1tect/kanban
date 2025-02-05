package eventBus

type Handler[T any] func(event T)

type EventBus[T any] struct {
	handlers map[string][]Handler[T]
}

func NewEventBus[T any]() *EventBus[T] {
	return &EventBus[T]{handlers: make(map[string][]Handler[T])}
}

func (eb *EventBus[T]) Subscribe(event string, handler Handler[T]) {
	eb.handlers[event] = append(eb.handlers[event], handler)
}

func (eb *EventBus[T]) Publish(event string, data T) {
	if handlers, found := eb.handlers[event]; found {
		for _, handler := range handlers {
			handler(data)
		}
	}
}
