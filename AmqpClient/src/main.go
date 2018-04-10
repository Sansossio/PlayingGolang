// Package
package main

// Imports
import (
	"fmt"
	"runtime"

	"./amqpclient"
)

// Properties
var (
	listeners     = runtime.NumCPU()
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
	fmt.Printf("AmqpClient on %s with architecture: %s over routines: %d\n", runtime.GOOS, runtime.GOARCH, listeners)
	// Properties
	amqpclient.SetProperties(listeners, msgs, printInterval)
	// Start listeners
	amqpclient.StartListeners(queues)
	// Scan
	fmt.Scanln()
}
