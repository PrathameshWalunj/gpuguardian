package monitor

import (
	"github.com/PrathameshWalunj/gpuguardian/internal/types"
)

// ProcessMonitor handles GPU process monitoring
type ProcessMonitor struct {
	processes chan []types.ProcessInfo
	stop      chan struct{}
}

// NewProcessMonitor creates a new process monitor
func NewProcessMonitor() *ProcessMonitor {
	return &ProcessMonitor{
		processes: make(chan []types.ProcessInfo, 100),
		stop:      make(chan struct{}),
	}
}

// Start begins the process monitoring
func (pm *ProcessMonitor) Start() error {
	// TODO
	return nil
}

// Stop ends the process monitoring
func (pm *ProcessMonitor) Stop() {
	close(pm.stop)
}

// Processes returns the channel for receiving process information
func (pm *ProcessMonitor) Processes() <-chan []types.ProcessInfo {
	return pm.processes
}
