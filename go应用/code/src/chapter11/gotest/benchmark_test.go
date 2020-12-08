package gotest

import (
	"fmt"
	"testing"
)

func Benchmark_Add(b *testing.B) {

	var n int
	for i := 0; i < b.N; i++ {
		n++
	}
}

func Benchmark_Add_TimerControl(b *testing.B) {

	b.ResetTimer()

	b.StartTimer()

	var n int
	for i := 0; i < b.N; i++ {
		n++
	}

	b.StopTimer()
}

func Benchmark_Alloc(b *testing.B) {

	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", i)
	}
}
