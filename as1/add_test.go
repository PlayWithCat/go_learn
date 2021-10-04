package test

import (
	"testing"
	"time"
)

func add(x, y int) int {
	return x + y
}

// func TestAdd(t *testing.T) {
// 	if add(1, 2) != 10 {
// 		t.Error()
// 	}

// 	time.Sleep(time.Second * 2)
// }

// func TestB(t *testing.T) {
// 	if add(1, 2) != 8 {
// 		t.Error()
// 	}

// 	time.Sleep(time.Second * 2)
// }

// func TestAdds(t *testing.T) {
// 	var tests = []struct {
// 		x      int
// 		y      int
// 		expect int
// 	}{
// 		{1, 1, 2},
// 		{2, 2, 0},
// 		{3, 2, 5},
// 	}

// 	for _, tt := range tests {
// 		actual := add(tt.x, tt.y)
// 		if actual != tt.expect {
// 			t.Errorf("add(%d, %d): expect %d, actual %d", tt.x, tt.y, tt.expect, actual)
// 		}
// 	}
// }

func BenchmarkAdd(b *testing.B) {
	println("B.N=", b.N)
	for i := 0; i < b.N; i++ {
		_ = add(1, 2)
	}
}

func sleep() {
	time.Sleep(time.Second * 1)
}

func sleep2() {
	time.Sleep(time.Second * 2)
}

func BenchmarkSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sleep()
		// sleep2()
	}
}

func BenchmarkSleep2(b *testing.B) {
	println("B.N=", b.N)
	sleep()
	// sleep2()
}
