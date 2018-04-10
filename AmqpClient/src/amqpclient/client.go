package amqpclient

import (
	"context"
	"log"
	"time"

	"pack.ag/amqp"
)

// Variables
var (
	connectedLog bool
	subscribeLog bool
)

// Send write message
func Send(consumer *Consumer) error {
	// Create a sender
	sender, err := consumer.session.NewSender(
		amqp.LinkTargetAddress("queue-name"),
	)
	// Error handler
	if err != nil {
		return err
	}
	// Context and time
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	// Defer
	defer sender.Close()
	defer cancel()
	// Send message
	err = sender.Send(ctx, amqp.NewMessage([]byte("Hello World!")))
	if err != nil {
		return err
	}
	// Response
	return nil
}

// Subscribe listener queue
func subscribe(queue string, consumer *Consumer) {
	// Create a receiver
	receiver, err := consumer.session.NewReceiver(
		amqp.LinkSourceAddress(queue),
		amqp.LinkCredit(1000),
	)
	// Comprobe error
	if err != nil {
		log.Fatal(err)
		return
	}
	// Msg handler
	// Channel
	ch := make(chan error)
	// Handler
	go loopHandler(receiver, queue, ch)
	// Callback
	go eventCallback("SUBSCRIBED")
	// Properties
	handlerErr := <-ch
	// Comprobe
	if handlerErr != nil {
		// Log
		log.Fatal(handlerErr)
	}
}

// Handler
func loopHandler(receiver *amqp.Receiver, topic string, channel chan error) {
	// Infinite loop
	for {
		// Receive next message
		msg, err := receiver.Receive(ctx)
		// Comprobe error
		if err != nil {
			channel <- err
			return
		}
		// Accept message
		msg.Accept()
		// Parse data
		msgToString := string(msg.GetData())
		// Handler
		go messageCallback(topic, msgToString)
	}
}

// Connect Create connection with amqp server
func connect(consumer *Consumer) error {
	// Create client
	client, err := amqp.Dial("amqp://localhost",
		amqp.ConnSASLPlain("access-key-name", "access-key"),
		amqp.ConnIdleTimeout(30*time.Hour),
	)
	if err != nil {
		return err
	}
	// Open a session
	session, err := client.NewSession()
	if err != nil {
		return err
	}
	// Create context
	if ctx == nil {
		ctx = context.Background()
	}
	// Consumer
	consumer.session = session
	consumer.client = client
	// Callback
	go eventCallback("CONNECTED")
	// Return
	return nil
}
