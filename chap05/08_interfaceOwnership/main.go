package main

import "github.com/disharjayanth/ultimate_go_notebook/chap05/08_interfaceOwnership/pubsub"

type Publisher interface {
	Publish(key string, value interface{}) error
	Subscribe(key string, value interface{}) error
}

type mock struct{}

func (m *mock) Publish(key string, value interface{}) error {
	// PRETEND THERE IS A SPECIFIC IMPLEMENTATION.
	return nil
}

func (m *mock) Subscribe(key string, value interface{}) error {
	// PRETEND THERE IS A SPECIFIC IMPLEMENTATION.
	return nil
}

func main() {
	pubs := []Publisher{
		&mock{},
		pubsub.New("localhost:8000"),
	}

	for _, p := range pubs {
		p.Publish("key", "value")
		p.Subscribe("key", "value")
	}
}
