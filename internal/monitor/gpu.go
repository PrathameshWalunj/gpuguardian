package monitor

import (
	"fmt"
	"log"
	"time"

	"github.com/PrathameshWalunj/gpuguardian/internal/types"
	"github.com/PrathameshWalunj/gpuguardian/pkg/nvml"
)

// Monitor is responsible for collecting GPU metrics periodically
// It abstracts the NVML library and provides an interface for real time GPU monitoring
type Monitor struct {
	updateInterval time.Duration         // Interval between successive GPU metric collections
	metrics        chan types.GPUMetrics // Channel for transmitting GPU metrics to the caller
	stop           chan struct{}         // Channel to signal the monitoring loop to stop
}

// NewMonitor initializes and returns a new Monitor instance with the specified update interval
func NewMonitor(updateInterval time.Duration) *Monitor {
	return &Monitor{
		updateInterval: updateInterval,
		metrics:        make(chan types.GPUMetrics, 100), // Buffered channel for efficient metric transmission
		stop:           make(chan struct{}),              // Channel used to stop the monitoring loop
	}
}

// Start initializes NVML and begins the GPU monitoring loop
// The metrics are periodically collected and sent to the `metrics` channel
func (m *Monitor) Start() error {
	if err := nvml.Init(); err != nil {
		return fmt.Errorf("failed to initialize NVML: %v", err)
	}

	go m.monitorLoop() // Launch the monitoring loop as a separate goroutine.
	return nil
}

// Stop terminates the monitoring loop and shuts down NVML
func (m *Monitor) Stop() {
	close(m.stop)   // Signal the monitoring loop to stop
	nvml.Shutdown() // Properly shut down NVML
}

// Metrics returns a read-only channel for accessing GPU metrics
func (m *Monitor) Metrics() <-chan types.GPUMetrics {
	return m.metrics
}

// monitorLoop continuously collects GPU metrics at the specified interval
// Metrics are sent to the `metrics` channel until the loop is stopped
func (m *Monitor) monitorLoop() {
	ticker := time.NewTicker(m.updateInterval) // Schedule periodic metric collection
	defer ticker.Stop()                        // Ensure the ticker is stopped when the loop ends

	for {
		select {
		case <-m.stop: // Stop signal received; exit the loop
			return
		case <-ticker.C: // Time to collect metrics
			metrics, err := m.collectMetrics()
			if err != nil {
				log.Printf("Error collecting metrics: %v", err) // Log errors but continue monitoring
				continue
			}
			m.metrics <- metrics // Send collected metrics to the channel
		}
	}
}

// collectMetrics gathers GPU metrics for the specified GPU index (default is 0)
func (m *Monitor) collectMetrics() (types.GPUMetrics, error) {
	return nvml.GetDeviceMetrics(0) // Collect metrics for the first GPU in the system
}
