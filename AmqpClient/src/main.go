// Package
package main

// Imports
import (
	"encoding/json"
	"fmt"
	"runtime"
	"strconv"

	"./amqpclient"
	"./results"
)

// Struct
type myJSON struct {
	Number int
}

// Properties
var (
	subscribedLog, connectedLog bool
	messageAsync                = true
	receiveJSON                 = true
	listeners                   = runtime.NumCPU()
	msgs                        = 300 * 1000
	printInterval               = 50000
	queues                      = []string{
		"Consumer.go.VirtualTopic.queue-one",
		"Consumer.go.VirtualTopic.queue-two",
	}
)

// Message callback
func onMessage(topic string, msg []byte) {
	// Properties
	value := 0
	// Manage messages here
	// Parse json
	if receiveJSON {
		// Json
		data := myJSON{}
		// Parse data
		json.Unmarshal(msg, &data)
		// Value
		value = data.Number
	} else {
		// Direct convertion
		message := string(msg)
		// Parse data
		msgToInt, _ := strconv.Atoi(message)
		// Value
		value = msgToInt
	}
	// Callback
	go results.PrintTime(value)
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
	/* JSONmsg:
	{"TIPO":"INFO","IDTIPO":"1021","ID":"002161","HD":"16922","HDL":" 39.2","TC":"80000","TCL":" 88","P":"1" ,"UE":"2013-05-18 12:08:22","FUE":"Destacado_BCNT1_A6AdvancedEdition_1920x1080.mov" ,"MUS":"S","CMP3":" 0" ,"CMM":"2","AMP3":"0" ,"CPU":"55","MEM":" 39.2" ,"CANAL":"5","DBM":" -200" ,"TDISP":"10","TS":"2017-02-15 05:00:00"}
	*/
	// Log
	fmt.Printf("AmqpClient on %s with architecture: %s over routines: %d\n", runtime.GOOS, runtime.GOARCH, listeners)
	// Liteners
	startListeners()
	// Scan
	fmt.Scanln()
}
