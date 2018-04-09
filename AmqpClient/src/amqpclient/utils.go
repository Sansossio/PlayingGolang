package amqpclient

import (
	"fmt"
	"time"
)

// Print time
func printTime(proccesed int) {
	// Properties
	// Execution time
	interval := time.Since(start)
	// Reset event
	start = time.Now()
	// Print
	fmt.Printf("Processed: %d, time: %s\n", printInterval, interval)
	// Print in finish
	if proccesed == msgt {
		fmt.Printf("Total time: %s total processed: %d, routines: %d\n", time.Since(totalTime), msgt, total)
	}
}

// Reconnec
func reconnect(err error, consumer *Consumer) {
	// Print error
	fmt.Println(err)
	fmt.Println("Retrying...")
	// Sleep
	time.Sleep(1 * time.Second)
	// Listener
	listener(consumer)
}

// Listener Create new listener instance
func listener(consumer *Consumer) *Consumer {
	// Properties
	err := connect(consumer)
	// Error handler
	if err != nil {
		reconnect(err, consumer)
	}
	// Subscribe
	for _, val := range consumer.queue {
		go subscribe(val, consumer)
	}
	// Return
	return consumer
}

// CreateConsumer Create new empty consumer
func createConsumer(id int, queues []string) Consumer {
	// Aux
	var consumer Consumer
	// Add
	consumer.id = id
	consumer.queue = queues
	// Return
	return consumer
}

// StartListeners Start new listeners
func StartListeners(queues []string) {
	// Loop
	for i := 0; i < total; i++ {
		// Create new instance
		consumer := createConsumer(i, queues)
		// Go async
		listener(&consumer)
	}
}

// SetProperties Set need properties
func SetProperties(t, m, pt int) {
	// Properties
	total = t          // Total time
	msgt = m           // Total messages to process
	printInterval = pt // Print interval
}
