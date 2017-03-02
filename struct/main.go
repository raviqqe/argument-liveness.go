package main

import "github.com/raviqqe/argument-liveness.go/list"

func main() {
	l := *list.NewList(42, nil)

	for {
		l = *l.Rest()
	}
}
