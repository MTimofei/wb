package natstreams

import (
	"log"
	"time"

	"github.com/nats-io/stan.go"
	"github.com/wb/cmd/0L/internal/config"
	"github.com/wb/pkg/erro"
)

const (
	role = "sub-client"
)

const (
	ErrConnect     = "NATS-STREAM cen't connect "
	ErrClose       = "NATS-STREAM cen't close connect "
	ErrSubscribe   = "NATS-STREAM cen't subscribe "
	ErrUnsubscribe = "NATS-STREAM cen't unsubscribe "
	EreWork        = "NATS-STREAM cen't work "
)

type NatsStreams struct {
	sc  stan.Conn
	sub stan.Subscription
}

func NewNatsStreams() (ns *NatsStreams, err error) {
	defer erro.IsError(ErrConnect, err)
	sc, err := stan.Connect(config.App.Cluster, role, stan.NatsURL(config.App.NatsURL))
	if err != nil {
		return nil, err
	}

	return &NatsStreams{sc: sc, sub: nil}, nil
}
func (ns *NatsStreams) Subscribe() (err error) {
	defer erro.IsError(ErrSubscribe, err)

	sub, err := ns.sc.Subscribe(
		config.App.Channel,
		handlerMsg,
		stan.DurableName("myApp"),
		stan.DeliverAllAvailable(),
		stan.SetManualAckMode(),
		stan.AckWait(2*time.Second),
	)
	if err != nil {
		return err
	}

	ns.sub = sub

	return nil
}
func (ns *NatsStreams) Work() (err error) {
	defer erro.IsError(EreWork, err)
	err = ns.Subscribe()
	if err != nil {
		return err
	}

	for {
		if ns.sub.IsValid() {
			time.Sleep(1 * time.Second)
		} else {
			log.Println("subject is not valid")
			break
		}
	}

	return nil
}

func (ns *NatsStreams) Unsubscribe() (err error) {
	defer erro.IsError(ErrUnsubscribe, err)

	err = ns.sub.Unsubscribe()
	if err != nil {
		return err
	}

	return nil
}

func (ns *NatsStreams) Close() (err error) {
	defer erro.IsError(ErrClose, err)

	err = ns.sc.Close()
	if err != nil {
		return err
	}
	return nil
}
