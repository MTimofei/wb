package natstreams

import (
	"fmt"
	"time"

	"github.com/nats-io/stan.go"
	"github.com/wb/cmd/0L/internal/config"
	"github.com/wb/pkg/erro"
)

const (
	role = "sub-client"
)

const (
	ErrInit              = " cen't init nats:"
	ErrClose             = " cen't close connect "
	ErrSubscribe         = " cen't subscribe "
	ErrUnsubscribe       = " cen't unsubscribe "
	EreWork              = " cen't work "
	ErrSubscribeNotValid = " subject is not valid "
)

type NatsStreams struct {
	sc        stan.Conn
	sub       stan.Subscription
	transport chan string
	ack       chan bool
}

func New(transport chan string, ack chan bool) (ns *NatsStreams, err error) {
	defer func(error) { erro.IsError(ErrInit, err) }(err)

	sc, err := stan.Connect(
		config.App.NS.Cluster,
		role,
		stan.NatsURL(config.App.NS.NatsURL),
	)
	if err != nil {
		return nil, err
	}

	fmt.Println("connect to: ", config.App.NS.NatsURL)
	return &NatsStreams{sc: sc, sub: nil, transport: transport, ack: ack}, nil
}
func (ns *NatsStreams) Subscribe() (err error) {
	defer func(error) { erro.IsError(ErrSubscribe, err) }(err)

	sub, err := ns.sc.Subscribe(
		config.App.NS.Channel,
		ns.handlerMsg,
		stan.DurableName("myApp"),
		stan.DeliverAllAvailable(),
		stan.SetManualAckMode(),
		stan.AckWait(2*time.Second),
	)
	if err != nil {
		return err
	}

	ns.sub = sub

	fmt.Println("subscribe to: ", config.App.NS.NatsURL)
	return nil
}
func (ns *NatsStreams) Work() (err error) {
	defer func(error) { erro.IsError(EreWork, err) }(err)

	err = ns.Subscribe()
	if err != nil {
		return err
	}
	fmt.Println("work nats")

	for {
		if ns.sub.IsValid() {
			time.Sleep(1 * time.Second)
		} else {
			err = fmt.Errorf("%s", ErrSubscribeNotValid)
			return err
		}
	}
}

func (ns *NatsStreams) Unsubscribe() (err error) {
	defer func(error) { erro.IsError(ErrUnsubscribe, err) }(err)

	err = ns.sub.Unsubscribe()
	if err != nil {
		return err
	}
	fmt.Println("unsubscribe nats")
	return nil
}

func (ns *NatsStreams) Close() (err error) {
	defer func(error) { erro.IsError(ErrClose, err) }(err)

	err = ns.sc.Close()
	if err != nil {
		return err
	}

	fmt.Println("close nats")
	return nil
}

func (ns *NatsStreams) Started() (err error) {
	defer func(error) { erro.IsError("NATS-STREAMS", err) }(err)
	// defer func(error) { err = ns.Close() }(err)
	// defer func(error) { err = ns.Unsubscribe() }(err)
	err = ns.Work()
	if err != nil {
		return err
	}
	return nil
}
