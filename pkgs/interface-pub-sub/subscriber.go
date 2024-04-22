package pubsub

type Subscriber[T any] interface {
	Subscribe(route string, callback func(T) error) error
}
