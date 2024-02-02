// Package server implements the API server running in public network
package api_server

import "net/http"

type APIServer struct {
	// server is the underlying http server
	server *http.Server
}

// NewAPIServer creates a new API server
func NewAPIServer(addr string, handler http.Handler) *APIServer {
	// create a new http server
	s := &http.Server{
		Addr:    addr,
		Handler: handler,
	}
	return &APIServer{server: s}
}
