package broker

type Broker interface {
	// chan string
	// chan os.Signal
	Started() error
	Subscribe() error
	Work() error
	Unsubscribe() error
	Close() error
}
