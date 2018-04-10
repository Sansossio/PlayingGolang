// Package
package main

// Imports
import (
	"fmt"
	"runtime"
	"strconv"

	"./amqpclient"
	"./results"
)

// Properties
var (
	subscribedLog, connectedLog bool
	messageAsync                = true
	listeners                   = runtime.NumCPU()
	msgs                        = 300000
	printInterval               = 50000
	queues                      = []string{
		"Consumer.go.VirtualTopic.queue-one",
		"Consumer.go.VirtualTopic.queue-two",
	}
)

// Message callback
func onMessage(topic, message string) {
	// Manage messages here
	// Parse data
	msgToInt, _ := strconv.Atoi(message)
	// Callback
	go results.PrintTime(msgToInt)
}

// Event callback
func onEvent(event string) {
	// Comprobe
	switch event {
	case "CONNECTED": // OnConnected event
		// Comprobe
		if !connectedLog {
			// Print
			fmt.Printf("Connected to amqp server\n")
			// Change value
			connectedLog = true
		}
		break
	case "SUBSCRIBED": // OnSubscribed event
		// Times
		results.Init()
		// Comprobe
		if !subscribedLog {
			// Print
			fmt.Printf("Subscribed to topics\n")
			// Change value
			subscribedLog = true
		}
		break
	}
}

// Amqp listeners
func startListeners() {
	// Properties
	results.SetProperties(msgs, printInterval)
	// Start listeners
	amqpclient.StartListeners(queues, listeners, onMessage, onEvent, messageAsync)
}

// Main
func main() {
	// Log
	fmt.Printf("AmqpClient on %s with architecture: %s over routines: %d\n", runtime.GOOS, runtime.GOARCH, listeners)
	// Liteners
	startListeners()
	// Scan
	fmt.Scanln()
}
