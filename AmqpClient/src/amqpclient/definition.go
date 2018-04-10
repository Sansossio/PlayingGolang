package amqpclient

import (
	"context"

	"pack.ag/amqp"
)

// Consumer Consumer of queue's
type Consumer struct {
	queue   []string
	client  *amqp.Client
	session *amqp.Session
}

// Global properties
var (
	ctx             context.Context
	messageAsync bool
	messageCallback func(string, string)
	eventCallback   func(string)
)
