package amqpclient

import (
	"context"
	"time"

	"pack.ag/amqp"
)

// Consumer Consumer of queue's
type Consumer struct {
	id      int
	queue   []string
	client  *amqp.Client
	session *amqp.Session
}

// Global properties
var (
	total, msgt, printInterval int
	totalTime, start           time.Time
	printed                    bool
	ctx                        context.Context
	firtsMsg                   bool
)
