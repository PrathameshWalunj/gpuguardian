package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PrathameshWalunj/gpuguardian/internal/monitor"
)

// main is the entry point of the GPU Guardian application
// It initializes the monitoring system and handles shutdown
func main() {

	log.SetFlags(log.Ltime | log.Lshortfile)

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

			fmt.Printf("\033[2J\033[H") // Clear screen
			fmt.Println("╔══════════════════════════════════════════╗")
			fmt.Println("║           GPU GUARDIAN MONITOR           ║")
			fmt.Println("╚══════════════════════════════════════════╝")
			fmt.Printf("\nGPU: %s\n", metrics.Name)
			fmt.Printf("Memory Usage: %.2f/%.2f GB (%.1f%%)\n",
				float64(metrics.MemoryUsed)/1024/1024/1024,
				float64(metrics.MemoryTotal)/1024/1024/1024,
				float64(metrics.MemoryUsed*100)/float64(metrics.MemoryTotal))
			fmt.Printf("GPU Utilization: %d%%\n", metrics.Utilization)
			fmt.Printf("Temperature: %d°C\n", metrics.Temperature)
			fmt.Printf("Process Count: %d\n\n", metrics.ProcessCount)
			fmt.Println("Press Ctrl+C to exit")
		}
	}()

	<-sigChan // Block until a signal is received
	fmt.Println("\nShutting down...")
}
