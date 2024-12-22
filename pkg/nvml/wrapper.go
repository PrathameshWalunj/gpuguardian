package nvml

/*
#cgo CFLAGS: -I"C:/Program Files/NVIDIA GPU Computing Toolkit/CUDA/v12.6/include"
#cgo windows LDFLAGS: -L"C:/Windows/System32" -lnvml
#include <nvml.h>
*/
import "C"
import (
	"fmt"

	"github.com/PrathameshWalunj/gpuguardian/internal/types"
)

// Init initializes the NVIDIA Management Library
// This function must be called before any other NVML functions
func Init() error {
	result := C.nvmlInit()
	if result != C.NVML_SUCCESS {
		return fmt.Errorf("NVML init failed: %v", result)
	}
	return nil
}

// Shutdown properly shuts down NVML and cleans up resources
func Shutdown() error {
	result := C.nvmlShutdown()
	if result != C.NVML_SUCCESS {
		return fmt.Errorf("NVML shutdown failed: %v", result)
	}
	return nil
}

// GetDeviceMetrics retrieves GPU metrics for the specified GPU index
func GetDeviceMetrics(index int) (types.GPUMetrics, error) {
	var device C.nvmlDevice_t
	var metrics types.GPUMetrics

	result := C.nvmlDeviceGetHandleByIndex(C.uint(index), &device)
	if result != C.NVML_SUCCESS {
		return metrics, fmt.Errorf("failed to get device handle: %v", result)
	}

	// Remaining metric collection omitted for brevity; identical to earlier code
	return metrics, nil
}
