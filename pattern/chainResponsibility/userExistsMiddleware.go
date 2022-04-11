package main

import "fmt"

type userExistsMiddleware struct {
	server *Server
	next   Middleware
}

func NewUserExistsMiddleware(server *Server) *userExistsMiddleware {
	return &userExistsMiddleware{
		server: server,
	}
}

func (u *userExistsMiddleware) check(email, password string) bool {
	if !u.server.hasEmail(email) {
		fmt.Println("This email is not registered")
		return false
	}

	if !u.server.isValidPassword(email, password) {
		fmt.Println("Wrong password!")
		return false
	}

	return u.checkNext(email, password)
}

func (u *userExistsMiddleware) checkNext(email, password string) bool {
	if u.next == nil {
		return true
	}

	return u.next.check(email, password)
}

func (u *userExistsMiddleware) linkWith(middleware Middleware) {
	u.next = middleware
}
