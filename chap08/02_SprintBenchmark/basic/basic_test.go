package basic

import (
	"fmt"
	"testing"
)

var gs string

func BenchmarkSprint(b *testing.B) {
	var s string
	// fmt.Println("Number of times test ran:", b.N)
	for i := 0; i < b.N; i++ {
		s = fmt.Sprint("hello")
	}
	gs = s
}

func BenchmarkSprintf(b *testing.B) {
	var s string
	// fmt.Println("Number of times test ran:", b.N)
	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("hello")
	}
	gs = s
}
