package sampl_test

import (
	"fmt"
	"testing"
)

var gs string

func benchsprint(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = fmt.Sprint("hello")
	}
	gs = s
}

func benchsprintf(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("hello")
	}
	gs = s
}

func BenchmarkSprint(b *testing.B) {
	b.Run("none", benchsprint)
	b.Run("format", benchsprintf)
}
