package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PrathameshWalunj/gpuguardian/internal/monitor"
)

// main is the entry point of the GPU Guardian application
// It initializes the monitoring system and handles shutdown
func main() {
	mon := monitor.NewMonitor(time.Second) // Create a new monitor with 1 second update interval

	if err := mon.Start(); err != nil {
		fmt.Printf("Failed to start monitoring: %v\n", err)
		os.Exit(1)
	}
	defer mon.Stop() // Ensure proper cleanup on exit

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM) // Handle interrupt signals

	go func() {
		for metrics := range mon.Metrics() {
			fmt.Printf("GPU Metrics: %+v\n", metrics) // Print metrics in real time
		}
	}()

	<-sigChan // Block until a signal is received
}
