package broker

type Broker interface {
	Subscribe() error
	Work() error
	Unsubscribe() error
	Close() error
}
