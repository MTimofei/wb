package broker

type Broker interface {
	Started() error
	Subscribe() error
	Work() error
	Unsubscribe() error
	Close() error
}
