package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		time.Sleep(time.Second * 3)
		// defer close(done)
		defer func() {
			done <- struct{}{}
		}()
		defer println("chan close")

		for {
			x, ok := <-c
			fmt.Printf("%v\n", ok)
			if !ok {
				return
			}

			println(x)
			println("continuing")
		}
	}()

	c <- 1
	c <- 2
	c <- 3

	close(c)
	<-done
}
