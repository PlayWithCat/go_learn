package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["1"] = 1
	m["2"] = 2

	fmt.Println(m)

	m2 := map[int]struct {
		x int
	}{
		1: {x: 100},
		2: {x: 200},
	}

	fmt.Println(m2)
	// delete(m, "1")
	// fmt.Printf("after delete: %+v\n", m)

	println(string('a' + 1))

	// var tmp int = m["1"]
	m["1"]++
	println("**************************")
	println(m["1"])
	// println(tmp)

	println(map[string]int{} == nil)
}
