package results

import (
	"fmt"
	"time"
)

// SetProperties Set need properties
func SetProperties(m, pt int) {
	// Properties
	totalMsg, printInterval = m, pt
}

// PrintTime Print proccess time
func PrintTime(proccesed int) {
	// Comprobation
	if proccesed%printInterval != 0 {
		return
	}
	// Execution time
	interval := time.Since(start)
	// Reset event
	start = time.Now()
	// Print
	fmt.Printf("Processed: %d, time: %s\n", printInterval, interval)
	// Print in finish
	if proccesed == totalMsg {
		fmt.Printf("Total time: %s total processed: %d\n", time.Since(totalTime), totalMsg)
	}
}

// Init start system
func Init() {
	start = time.Now()
	totalTime = time.Now()
}
