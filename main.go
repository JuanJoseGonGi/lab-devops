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

func (s Server) StartFileServer() error {
	s.Handler = http.FileServer(http.Dir("./static"))
	return s.Start()
}

func main() {
	server := NewServer(":8080")

	if err := server.StartFileServer(); err != nil {
		panic(err)
	}
}
