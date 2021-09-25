package main

import (
	"sync"
	"time"
)

type data struct {
	sync.Mutex
}

func (d data) test(s string) {
	d.Lock()
	defer d.Unlock()

	for i := 0; i < 5; i++ {
		println(s, i)
		time.Sleep(time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	var d data
	p := &d

	go func() {
		defer wg.Done()
		p.test("read")
	}()

	go func() {
		defer wg.Done()
		p.test("write")
	}()

	wg.Wait()
}
