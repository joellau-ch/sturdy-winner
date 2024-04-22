package pubsub

type PubSubber[T any] interface {
	Publisher[T]
	Subscriber[T]
}

type PubSubberMessage[T any] any
