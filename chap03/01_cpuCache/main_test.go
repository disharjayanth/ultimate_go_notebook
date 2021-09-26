package main

import "testing"

func BenchmarkRowTraverse(b *testing.B) {
	b.Run("Row Traverse", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			RowTraverse()
		}
	})
}

func BenchmarkColTraverse(b *testing.B) {
	b.Run("Col Traverse", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ColTraverse()
		}
	})
}

func BenchmarkLinkedListTraverse(b *testing.B) {
	b.Run("LinkedList Traverse", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			LinkedListTraverse()
		}
	})
}
