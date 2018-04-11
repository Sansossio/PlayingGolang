package amqpclient

import (
	"fmt"
	"time"
)

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
	// Consumer
	consumer := Consumer{queue: queues}
	// Return
	return consumer
}

// StartListeners Start new listeners
func StartListeners(queues []string, total int, onMessage func(string, []byte), onEvent func(string), ca bool) {
	// Properties
	messageAsync, messageCallback, eventCallback = ca, onMessage, onEvent
	// Loop
	for i := 0; i < total; i++ {
		// Create new instance
		consumer := createConsumer(i, queues)
		// Go async
		listener(&consumer)
	}
}
