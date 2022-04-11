package main

import "fmt"

type Server struct {
	middleware Middleware
	users      map[string]string
}

func NewServer() *Server {
	return &Server{
		users: make(map[string]string),
	}
}

func (s *Server) setMiddleware(middleware Middleware) {
	s.middleware = middleware
}

func (s *Server) logIn(email, password string) bool {
	if s.middleware.check(email, password) {
		fmt.Println("Authorization has been successful!")
		return true
	}

	return false
}

func (s *Server) register(email, password string) {
	s.users[email] = password
}

func (s *Server) hasEmail(email string) bool {
	_, ok := s.users[email]
	return ok
}

func (s *Server) isValidPassword(email, password string) bool {
	if s.users[email] == password {
		return true
	}
	return false
}
