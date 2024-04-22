package pubsub

type Publisher[T any] interface {
	Publish(route string, message T) error
}
