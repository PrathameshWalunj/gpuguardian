package nvml

/*
#cgo CFLAGS: -I"C:/Program Files/NVIDIA GPU Computing Toolkit/CUDA/v12.6/include"
#cgo windows LDFLAGS: -L"C:/Windows/System32" -L"${SRCDIR}/../../libs" -lnvml -lucrt
#cgo windows LDFLAGS: -static
#cgo linux LDFLAGS: -lnvidia-ml
#include <nvml.h>
*/
import "C"
import (
	"fmt"
	"log"

	"github.com/PrathameshWalunj/gpuguardian/internal/types"
)

// Init initializes the NVIDIA Management Library
// This function must be called before any other NVML functions
func Init() error {
	log.Println("Initializing NVML...")
	result := C.nvmlInit()
	if result != C.NVML_SUCCESS {
		return fmt.Errorf("NVML init failed: %v", result)
	}
	log.Println("NVML initialized successfully")
	return nil
}

// Shutdown properly shuts down NVML and cleans up resources
func Shutdown() error {
	log.Println("Shutting down NVML...")
	result := C.nvmlShutdown()
	if result != C.NVML_SUCCESS {
		return fmt.Errorf("NVML shutdown failed: %v", result)
	}
	log.Println("NVML shut down successfully")
	return nil
}

// GetDeviceMetrics retrieves GPU metrics for the specified GPU index
func GetDeviceMetrics(index int) (types.GPUMetrics, error) {
	var device C.nvmlDevice_t
	var metrics types.GPUMetrics

	log.Printf("Getting device handle for index %d...\n", index)
	result := C.nvmlDeviceGetHandleByIndex(C.uint(index), &device)
	if result != C.NVML_SUCCESS {
		return metrics, fmt.Errorf("failed to get device handle: %v", result)
	}

	// Get device name
	var name [C.NVML_DEVICE_NAME_BUFFER_SIZE]C.char
	log.Println("Getting device name...")
	result = C.nvmlDeviceGetName(device, &name[0], C.NVML_DEVICE_NAME_BUFFER_SIZE)
	if result == C.NVML_SUCCESS {
		metrics.Name = C.GoString(&name[0])
		log.Printf("Device name: %s\n", metrics.Name)
	} else {
		log.Printf("Failed to get device name: %v\n", result)
	}

	// Get memory info
	var memory C.nvmlMemory_t
	log.Println("Getting memory info...")
	result = C.nvmlDeviceGetMemoryInfo(device, &memory)
	if result == C.NVML_SUCCESS {
		metrics.MemoryTotal = uint64(memory.total)
		metrics.MemoryUsed = uint64(memory.used)
		metrics.MemoryFree = uint64(memory.free)
		log.Printf("Memory - Total: %d, Used: %d, Free: %d\n",
			metrics.MemoryTotal, metrics.MemoryUsed, metrics.MemoryFree)
	} else {
		log.Printf("Failed to get memory info: %v\n", result)
	}

	// Get utilization rates
	var utilization C.nvmlUtilization_t
	log.Println("Getting utilization rates...")
	result = C.nvmlDeviceGetUtilizationRates(device, &utilization)
	if result == C.NVML_SUCCESS {
		metrics.Utilization = uint32(utilization.gpu)
		log.Printf("GPU utilization: %d%%\n", metrics.Utilization)
	} else {
		log.Printf("Failed to get utilization rates: %v\n", result)
	}

	// Get temperature
	var temp C.uint
	log.Println("Getting temperature...")
	result = C.nvmlDeviceGetTemperature(device, C.NVML_TEMPERATURE_GPU, &temp)
	if result == C.NVML_SUCCESS {
		metrics.Temperature = uint32(temp)
		log.Printf("Temperature: %d°C\n", metrics.Temperature)
	} else {
		log.Printf("Failed to get temperature: %v\n", result)
	}

	metrics.Index = index
	return metrics, nil
}
