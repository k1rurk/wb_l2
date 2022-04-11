package main

type Middleware interface {
	linkWith(Middleware)
	check(string, string) bool
	checkNext(string, string) bool
}
