package main

import (
	"fmt"
)

func main() {
	m := "ζθΏ"
	for i, c := range m {
		fmt.Printf("%d: [%c]", i, c)
	}

	r := 'ζ'

	s := string(r)
	b := byte(r)

	s2 := string(b)
	r2 := rune(b)

	fmt.Println(s, b, s2, r2)
}
