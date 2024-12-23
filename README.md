# GPU Guardian

**GPU Guardian** is a real time GPU monitoring and security tool that focuses on detecting cryptomining malware and analyzing GPU process behavior.

---
Current Output


![Current Output](https://github.com/user-attachments/assets/7a0e279f-6e2b-4d97-ac25-4d4188df4e36)

## Features

### Core Features
- **Real time GPU metrics monitoring**:
  - Memory usage tracking
  - Temperature monitoring
  - Utilization analysis
  - Process level GPU usage tracking

### Security Features (Soon)
- Cryptomining detection using behavioral analysis
- Suspicious GPU usage pattern detection
- Process isolation monitoring
- Real time alerts for abnormal behavior

### Dashboard & Visualization (In Development)
- Web based dashboard
- Real time performance graphs
- Process activity visualization
- Alert management

---

## Technical Requirements

### Prerequisites
- NVIDIA GPU with latest drivers
- CUDA Toolkit (12.x or later)
- Go 1.21 or later
- MinGW-w64 (for Windows)

### Supported GPUs
- NVIDIA GeForce Series
- NVIDIA Quadro Series
- Other NVIDIA GPUs with NVML support

---

## Installation

```bash
# Clone the repository
git clone https://github.com/PrathameshWalunj/gpuguardian.git

# Navigate to project directory
cd gpuguardian

# Install dependencies
go mod tidy

# Run the application
go run cmd/guardian/main.go
