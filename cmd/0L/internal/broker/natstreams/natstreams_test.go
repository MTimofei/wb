package natstreams_test

import (
	"github.com/wb/cmd/0L/internal/broker"
	"github.com/wb/cmd/0L/internal/broker/natstreams"
)

var _ broker.Broker = (*natstreams.NatsStreams)(nil)
