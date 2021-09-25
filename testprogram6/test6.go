package main

import (
	"fmt"
)

type data struct{}

type tester interface {
	test()
	string() string
}

func (d *data) test() {
	fmt.Printf("addr of *data in test: %p\n", d)
}

func (d data) string() string {
	fmt.Printf("addr of data in string: %p\n", &d)
	return "string"
}

func main() {
	var d data
	var t tester = d

	fmt.Printf("data: %p\n", &d)
	t.test()
	println(t.string())
}
