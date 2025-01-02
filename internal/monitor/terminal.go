package monitor

import (
	"fmt"

	"github.com/PrathameshWalunj/gpuguardian/internal/types"
)

// RunTerminalUI starts the terminal user interface
func RunTerminalUI(metrics <-chan types.GPUMetrics) {
	for metric := range metrics {
		fmt.Printf("\033[2J\033[H") // Clear screen
		fmt.Println("╔══════════════════════════════════════════╗")
		fmt.Println("║           GPU GUARDIAN MONITOR           ║")
		fmt.Println("╚══════════════════════════════════════════╝")
		fmt.Printf("\nGPU: %s\n", metric.Name)
		fmt.Printf("Memory Usage: %.2f/%.2f GB (%.1f%%)\n",
			float64(metric.MemoryUsed)/1024/1024/1024,
			float64(metric.MemoryTotal)/1024/1024/1024,
			float64(metric.MemoryUsed*100)/float64(metric.MemoryTotal))
		fmt.Printf("GPU Utilization: %d%%\n", metric.Utilization)
		fmt.Printf("Temperature: %d°C\n", metric.Temperature)
		fmt.Printf("Process Count: %d\n\n", metric.ProcessCount)
		fmt.Println("Press Ctrl+C to exit")
	}
}
