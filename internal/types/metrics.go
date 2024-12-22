package types

// GPUMetrics represents the key GPU monitoring information that the application tracks
// It provides a summary of GPU usage, including memory, utilization, and other relevant stats
type GPUMetrics struct {
	Index        int    `json:"index"`         // GPU index as reported by NVML (0-based)
	Name         string `json:"name"`          // Human-readable name of the GPU
	MemoryTotal  uint64 `json:"memory_total"`  // Total GPU memory in bytes
	MemoryUsed   uint64 `json:"memory_used"`   // Currently used GPU memory in bytes
	MemoryFree   uint64 `json:"memory_free"`   // Free GPU memory in bytes
	Utilization  uint32 `json:"utilization"`   // Current GPU utilization as a percentage (0-100)
	Temperature  uint32 `json:"temperature"`   // Current GPU temperature in Celsius
	PowerUsage   uint32 `json:"power_usage"`   // Current GPU power usage in milliwatts
	ProcessCount uint32 `json:"process_count"` // Number of processes actively using the GPU
}

// ProcessInfo contains information about individual processes utilizing the GPU.
type ProcessInfo struct {
	PID         uint32 `json:"pid"`         // Process ID of the GPU user
	Name        string `json:"name"`        // Name of the process
	MemoryUsed  uint64 `json:"memory_used"` // GPU memory used by the process in bytes
	Utilization uint32 `json:"utilization"` // GPU utilization by the process as a percentage
}
