package main

import "github.com/raviqqe/argument-liveness.go/list"
import "time"

func main() {
	l := list.NewList(42, nil)

	for {
		l = l.Rest()
		time.Sleep(1000)
	}
}
