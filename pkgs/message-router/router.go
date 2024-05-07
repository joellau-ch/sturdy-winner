package simplemessagerouter

import (
	"log"

	"github.com/google/uuid"
)

type SimpleMessageRouter struct {
	routes map[RouteKey]Subscriptions
}

var _ MessageRouter[RouteKey, Message, SubscriptionId] = &SimpleMessageRouter{}

type RouteKey struct {
	Route string
	Type  string
}
type Message any
type SubscriptionId string
type Subscriptions map[SubscriptionId]HandlerFunc[Message] // need a sub-map for unsubscribing

func NewSimpleMessageRouter() (router *SimpleMessageRouter, err error) {
	router = &SimpleMessageRouter{}
	router.routes = map[RouteKey]Subscriptions{}

	return
}

func (m *SimpleMessageRouter) Publish(key RouteKey, message Message) (err error) {
	ensureKeyExists(key, m.routes)

	for _, handler := range m.routes[key] {
		err = handler(message)

		// TODO: add ability to retry automatically
		// TODO: move error handling somewhere to prevent
		//       blocking calls to other subscribers
		if err != nil {
			return
		}
	}

	// TODO: change error logic, errors should reflect problems in
	//       sending messages, not errors with consumers
	return
}

func (m *SimpleMessageRouter) Subscribe(key RouteKey, handler HandlerFunc[Message]) (id SubscriptionId, err error) {
	ensureKeyExists(key, m.routes)

	// generate unique id
	// consider using less expensive way of generating ids
	uuid, err := uuid.NewRandom()
	if err != nil {
		return
	}
	id = SubscriptionId(uuid.String())

	// register the handler
	m.routes[key][id] = handler
	return
}

func (m *SimpleMessageRouter) Unsubscribe(key RouteKey, subId SubscriptionId) error {
	ensureKeyExists(key, m.routes)
	if _, found := m.routes[key][subId]; !found {
		log.Printf("Subscription not found: %v\n", subId)
		return nil
	}

	delete(m.routes[key], subId)
	return nil
}

func ensureKeyExists(key RouteKey, routes map[RouteKey]Subscriptions) {
	_, found := routes[key]
	if !found {
		routes[key] = Subscriptions{}
	}
}
