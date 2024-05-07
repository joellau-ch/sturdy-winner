// supporting types
package simplemessagerouter

// T: topic key
//   - must allow allow comparison via equality operator
//
// M: messages type
//
// S: subscription key
//   - must allow allow comparison via equality operator
//   - used for unsubscribing
type MessageRouter[T any, M any, S any] interface {
	Publisher[T, M]
	Subscriber[T, M, S]
}

type Publisher[T any, M any] interface {
	// broadcasts the message to anyone who has subscribed to the topic
	Publish(topic T, message M) error
}

type Subscriber[T any, M any, S any] interface {
	// registers a new handler to the topic
	Subscribe(topic T, handler HandlerFunc[M]) (S, error)

	// removes subscriber from the topic
	Unsubscribe(topic T, subId S) error
}

type HandlerFunc[M any] func(M) error
