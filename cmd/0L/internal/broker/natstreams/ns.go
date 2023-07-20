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
	ErrConnect           = " cen't connect "
	ErrClose             = " cen't close connect "
	ErrSubscribe         = " cen't subscribe "
	ErrUnsubscribe       = " cen't unsubscribe "
	EreWork              = " cen't work "
	ErrSubscribeNotValid = " subject is not valid "
)

type NatsStreams struct {
	sc  stan.Conn
	sub stan.Subscription
}

func NewNatsStreams() (ns *NatsStreams, err error) {
	defer erro.IsError(ErrConnect, err)

	sc, err := stan.Connect(
		config.App.NS.Cluster,
		role,
		stan.NatsURL(config.App.NS.NatsURL),
	)
	if err != nil {
		return nil, err
	}

	return &NatsStreams{sc: sc, sub: nil}, nil
}
func (ns *NatsStreams) Subscribe() (err error) {
	defer erro.IsError(ErrSubscribe, err)

	sub, err := ns.sc.Subscribe(
		config.App.NS.Channel,
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
			err = fmt.Errorf("%s", ErrSubscribeNotValid)
			return err
		}
	}
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

func (ns *NatsStreams) Started() (err error) {
	defer erro.IsError("NATS-STREAMS", err)
	defer func(error) { err = ns.Close() }(err)
	defer func(error) { err = ns.Unsubscribe() }(err)
	err = ns.Work()
	if err != nil {
		return err
	}
	return nil
}
