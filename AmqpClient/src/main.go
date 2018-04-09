// Package
package main

// Imports
import (
	"fmt"

	"./amqpclient"
)

// Properties
var (
	listeners     = 8
	msgs          = 300000
	printInterval = 50000
	queues        = []string{
		"Consumer.go.VirtualTopic.queue-one",
		"Consumer.go.VirtualTopic.queue-two",
	}
)

// Main
func main() {
	// Log
	fmt.Printf("Starting system\n")
	// Properties
	amqpclient.SetProperties(listeners, msgs, printInterval)
	// Start listeners
	amqpclient.StartListeners(queues)
	// Scan
	fmt.Scanln()
}
