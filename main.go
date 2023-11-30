package main

import "net/http"

type Server struct {
	*http.Server
}

func NewServer(addr string) *Server {
	return &Server{
		Server: &http.Server{
			Addr:    addr,
			Handler: nil,
		},
	}
}

func (s Server) Start() error {
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	return nil
}

func launchFileServer(addr string) error {
	server := NewServer(addr)
	server.Handler = http.FileServer(http.Dir("./static"))

	return server.Start()
}

func main() {
	if err := launchFileServer(":8080"); err != nil {
		panic(err)
	}
}
