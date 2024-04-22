package appdcc

import interfacepubsub "github.com/coinhako/joellau-ch/sturdy-winner/pkgs/interface-pub-sub"

type SimpleMessageRouter[T any] struct {
	routes map[string]func(T) error
}

var _ interfacepubsub.PubSubber[any] = &SimpleMessageRouter[any]{}

func NewSimpleMessageRouter() (router *SimpleMessageRouter[any]) {
	return
}

func (d *SimpleMessageRouter[T]) Publish(route string, message T) (err error) {
	return d.routes[route](message)
}

func (d *SimpleMessageRouter[T]) Subscribe(route string, callback func(T) error) (err error) {
	d.routes[route] = callback
	return
}
