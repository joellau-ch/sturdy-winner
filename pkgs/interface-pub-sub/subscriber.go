package pubsub

type Subscriber[T any] interface {
	Subscribe(route string, callback Handler[T]) error
}

type Handler[T any] func(T) error
