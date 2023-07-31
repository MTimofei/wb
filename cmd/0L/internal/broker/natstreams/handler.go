package natstreams

import (
	"fmt"

	"github.com/nats-io/stan.go"
	"github.com/wb/cmd/0L/internal/inspector"
)

func (ns *NatsStreams) handlerMsg(msg *stan.Msg) {
	if inspector.Check(msg.Data) {
		ns.transport <- fmt.Sprintf("%s", msg.Data)

		// select {
		ok := <-ns.ack
		if ok {
			msg.Ack()
		} else {
			return
		}
	}
	msg.Ack()

}
