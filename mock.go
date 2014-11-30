package main

import (
	"fmt"
)

type presenter interface {
	method() interface{}
}

type doing func() interface{}

func (d doing) method() interface{} {
	return d()
}

func mocking(d doing, p presenter) presenter {
	if d == nil {
		return p
	}
	return d
}

type actor struct {
	name string
}

func (a *actor) method() interface{} {
	println("actor method")
	return "actor"
}

var mock = func() interface{} {
	println("mock method")
	return "mock"
}

func main() {
	var p presenter
	p = new(actor)
	fmt.Printf("%v",p.method())

	p = mocking(mock,p)
	fmt.Printf("%v",p.method())
}