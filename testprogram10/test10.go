package main

type N int

func (N) value() {
	println("value")
}
func (*N) pointer() {
	println("pointer")
}

func main() {
	var t N
	var p *N = &t
	// pp := &p

	p.pointer()
	p.value()
}
