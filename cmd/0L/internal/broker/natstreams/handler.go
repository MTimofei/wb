package natstreams

import (
	"fmt"

	"github.com/nats-io/stan.go"
)

func (ns *NatsStreams) handlerMsg(msg *stan.Msg) {
	// fmt.Println("handlerMsg")
	// fmt.Println(msg.RedeliveryCount)
	ns.transport <- fmt.Sprintf("%s", msg.Data)

	// select {
	ok := <-ns.ack
	if ok {
		msg.Ack()
	} else {
		return
	}
	// }
}
