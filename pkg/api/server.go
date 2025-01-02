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
func NewServer(metrics chan types.GPUMetrics) *Server {
	return &Server{
		metrics: metrics,
	}
}

func (s *Server) Start() error {
	// Serve static files
	http.Handle("/", http.FileServer(http.Dir("web/build")))

	// WebSocket endpoint for real-time metrics
	http.HandleFunc("/ws", s.handleWebSocket)

	// Start broadcasting metrics to all connected clients
	go s.broadcastMetrics()

	log.Printf("Starting server on :8080")
	return http.ListenAndServe(":8080", nil)
}

func (s *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}

	// Store new client connection
	clientID := fmt.Sprintf("%p", conn)
	s.clients.Store(clientID, conn)

	log.Printf("New client connected: %s", clientID)

	// Clean up on disconnect
	conn.SetCloseHandler(func(code int, text string) error {
		s.clients.Delete(clientID)
		log.Printf("Client disconnected: %s", clientID)
		return nil
	})
}
