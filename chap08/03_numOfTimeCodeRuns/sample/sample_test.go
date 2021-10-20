package sample_test

import (
	"fmt"
	"testing"
)

var gs string
var a []int

func BenchmarkSprint(b *testing.B) {
	var s string

	a = append(a, b.N)
	for i := 0; i < b.N; i++ {
		s = fmt.Sprint("Hello")
	}

	if len(a) > 4 {
		fmt.Println(a)
	}

	gs = s
}

func BenchmarkSprintf(b *testing.B) {
	var s string

	a = append(a, b.N)
	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("Hello")
	}

	if len(a) > 4 {
		fmt.Println(a)
	}

	gs = s
}
