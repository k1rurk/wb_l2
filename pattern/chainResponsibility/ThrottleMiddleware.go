package main

import (
	"fmt"
	"time"
)

type ThrottleMiddleware struct {
	next Middleware
	requestPerMinute,
	request int
	currentTime int64
}

func NewThrottleMiddleware(requestPerMinute int) *ThrottleMiddleware {
	return &ThrottleMiddleware{
		requestPerMinute: requestPerMinute,
		currentTime:      time.Now().UnixNano(),
	}
}

func (t *ThrottleMiddleware) check(email, password string) bool {
	if time.Now().UnixNano() > t.currentTime+60_000 {
		t.request = 0
		t.currentTime = time.Now().UnixNano()
	}

	t.request++

	if t.request > t.requestPerMinute {
		fmt.Println("Request limit exceeded")
		return false
	}

	return t.checkNext(email, password)
}

func (t *ThrottleMiddleware) checkNext(email, password string) bool {
	if t.next == nil {
		return true
	}

	return t.next.check(email, password)
}

func (t *ThrottleMiddleware) linkWith(middleware Middleware) {
	t.next = middleware
}
