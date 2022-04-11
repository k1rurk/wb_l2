package main

import "fmt"

type roleCheckMiddleware struct {
	next Middleware
}

func (r *roleCheckMiddleware) check(email, password string) bool {
	if email == "admin@example.com" {
		fmt.Println("Hello, admin!")
		return true
	}

	fmt.Println("Hello, user!")
	return r.checkNext(email, password)
}

func (r *roleCheckMiddleware) checkNext(email, password string) bool {
	if r.next == nil {
		return true
	}

	return r.next.check(email, password)
}

func (r *roleCheckMiddleware) linkWith(middleware Middleware) {
	r.next = middleware
}
