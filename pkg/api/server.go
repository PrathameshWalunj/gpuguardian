package api

import (
	"fmt"
	"log"
	"net/http" //to create HTTP server and handle HTTP requests
	"sync"     // concurrent safe data structure
	"time"

	"github.com/PrathameshWalunj/gpuguardian/internal/types"
	"github.com/gorilla/websocket"
)

/*
* upgrader is a WebSocket upggrader that upgrades HTTP connections to WebSocket Connecrtions
* CheckOrigin allows all connections for dev purposes
 */
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WebSocket Server that broadcasts GPU metrics
type Server struct {

	// channel to
	metrics chan types.GPUMetrics
	clients sync.Map
}
