package main

import "fmt"

type Server struct {
	Port int
	Host string
	TLS  bool
}

func NewPort(n int) Option {
	return func(s *Server) {
		s.Port = n
	}
}

func NewHost(n string) Option {
	return func(s *Server) {
		s.Host = n
	}
}

func NewTLS(n bool) Option {
	return func(s *Server) {
		s.TLS = n
	}
}

type Option func(*Server)

func NewServer(opts ...Option) *Server {
	s := &Server{
		1234,
		"Hostlog",
		true,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func main() {
	s1 := NewServer()
	fmt.Println(s1)

	s2 := NewServer(
		NewPort(4321),
		NewHost("google.com"),
		NewTLS(false))
	fmt.Println(s2)
}
