package main

import (
	"fmt"
	"os"
)

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

var server *Server

func init() {
	server = NewServer()
	server.register("admin@example.com", "admin123")
	server.register("user@example.com", "user123")

	tMiddleware := NewThrottleMiddleware(2)
	uEMiddleware := NewUserExistsMiddleware(server)
	roleMiddleware := &roleCheckMiddleware{}
	uEMiddleware.linkWith(roleMiddleware)
	tMiddleware.linkWith(uEMiddleware)

	server.setMiddleware(tMiddleware)
}

func main() {

	success := false

	for !success {
		fmt.Println("Enter email: ")
		var email, password string
		fmt.Fscan(os.Stdin, &email)
		fmt.Println("Enter password: ")
		fmt.Fscan(os.Stdin, &password)
		success = server.logIn(email, password)
	}

}
