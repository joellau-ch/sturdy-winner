package pubsub

type PubSubber[T any] interface {
	// sends messages to a specific route
	Publisher[T]

	// register callbacks that are triggered on matching route
	Subscriber[T]
}
