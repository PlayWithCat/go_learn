package main

func test(x int) func() {
	println(&x)
	return func() {
		println(&x, x)
	}
}

func test1() []func() {
	var s []func()

	for i := 0; i < 2; i++ {
		x := i
		s = append(s, func() {
			println(&x, x)
		})
	}

	return s
}

func main() {
	// f := test(0x100)
	// f()

	for _, f := range test1() {
		f()
	}
}
