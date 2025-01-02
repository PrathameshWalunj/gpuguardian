package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PrathameshWalunj/gpuguardian/internal/monitor"
	"github.com/PrathameshWalunj/gpuguardian/pkg/api"
)

func main() {
	// Command line flags for configuring the application's behavior
	// 'web' flag determines if the app runs with a web dashboard
	// 'port' flag sets the port for the web server
	webMode := flag.Bool("web", false, "Run in web mode with dashboard")
	port := flag.String("port", "8080", "Port for web server")
	flag.Parse() // Parse the command-line flags

	// Set up the monitoring system with a 1-second update interval
	mon := monitor.NewMonitor(time.Second)
	if err := mon.Start(); err != nil {
		// If the monitoring system fails to start, log the error and exit
		log.Fatalf("Failed to start monitoring: %v\n", err)
	}
	defer mon.Stop() // Ensure the monitor stops cleanly when the program exits

	// Set up signal handling for graceful shutdown (e.g., SIGINT, SIGTERM)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	if *webMode {
		// If webMode is true, start the web server to serve the dashboard
		server := api.NewServer(mon.Metrics())
		go func() {
			// Start the web server in a separate goroutine
			if err := server.Start(); err != nil {
				// If the server fails to start, log the error and exit
				log.Fatalf("Server failed: %v\n", err)
			}
		}()
		log.Printf("Web dashboard running at http://localhost:%s\n", *port)

		// Wait for shutdown signal
		<-sigChan
	} else {
		// If webMode is false, run in terminal mode to display metrics in the console
		go monitor.RunTerminalUI(mon.Metrics())
		// Wait for shutdown signal
		<-sigChan
	}

	// shut down the application after receiving a signal
	log.Println("Shutting down...")
}
