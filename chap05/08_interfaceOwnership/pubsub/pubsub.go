package pubsub

type PubSub struct {
	host string
}

func New(host string) *PubSub {
	return &PubSub{
		host: host,
	}
}

func (ps *PubSub) Publish(key string, value interface{}) error {
	// PRETEND THERE IS A SPECIFIC IMPLEMENTATION.
	return nil
}

func (ps *PubSub) Subscribe(key string, value interface{}) error {
	// PRETEND THERE IS A SPECIFIC IMPLEMENTATION.
	return nil
}
