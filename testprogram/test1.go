package main

import (
	"fmt"
	"log"
	"math"
	"time"

	_ "pkgt"
)

type user struct {
	name string
	age  byte
}

func (u user) ToString() string {
	return fmt.Sprintf("%+v", u)
}

type manager struct {
	user
	title string
}

func test() *int {
	a := 0x100
	return &a
}

type serverOption struct {
	address string
	port    int
	path    string
	timeout time.Duration
	log     *log.Logger
}

func newOperion() *serverOption {
	return &serverOption{
		address: "0.0.0.0",
		port:    8080,
		path:    "var/test",
		timeout: time.Second * 5,
		log:     nil,
	}
}

func server(option *serverOption) {
	fmt.Printf("%+v", option)
}

func return2values() (int, int) {
	return 1, 1
}

func add() (z int) {
	{
		z := 10
		return z
	}
}

func main() {

	func(s string) { println(s) }("hello hello")

	println(add())

	p := add
	println(p())

	opt := newOperion()
	opt.port = 8050
	server(opt)

	println("hello world!!")
	println("this is printed by println")

	x := []int{100, 101, 102}
	for i, n := range x {
		println(i, ":", n)
	}

	var m manager
	m.name = "Tom"
	m.age = 29
	m.title = "Tomas"

	println(m.ToString())

	a := 100
	println(&a, a)
	{
		a := 200
		println(&a, a)
	}
	println(&a, a)

	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
	)
	println(KB, MB, GB)

	println(math.MaxInt16, math.MinInt16)

	const v = 20
	var aa byte = 10
	bb := v + aa
	fmt.Printf("%T, %v\n", bb, bb)

	tt := test()
	println(tt, *tt)

	println(return2values())

	for i := 0; i < 3; i++ {
		x := i
		println(&x, x)
	}
}
