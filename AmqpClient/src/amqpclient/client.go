package amqpclient

import (
	"context"
	"fmt"
	"log"
	"strconv"
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
	go loopHandler(receiver, ch)
	// Log
	if !subscribeLog {
		// Properties
		subscribeLog = true
		// Control time
		start = time.Now()
		totalTime = time.Now()
		// Print
		fmt.Printf("Subscribed to topics\n")
	}
	// Properties
	handlerErr := <-ch
	// Comprobe
	if handlerErr != nil {
		// Log
		log.Fatal(handlerErr)
	}
}

// Handler
func loopHandler(receiver *amqp.Receiver, channel chan error) {
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
		// Handler
		go msgHandler(msg)
	}
}

// MsgHandler
func msgHandler(msg *amqp.Message) {
	// Parse data
	msgToString := string(msg.GetData())
	msgToInt, _ := strconv.Atoi(msgToString)
	// Validate
	if msgToInt%printInterval == 0 {
		// Print time
		printTime(msgToInt) // Print
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
	// Log
	if !connectedLog {
		// Print
		fmt.Printf("Connected to amqp server\n")
		// Properties
		connectedLog = true
	}
	// Return
	return nil
}
