package testing

import (
	"math/rand"
	"testing"
	"time"
)

//RunSortBenchmarks runs sorting benchmarks using the given sorting func
func RunSortBenchmarks(b *testing.B, sort func([]int) []int) {

	b.ReportAllocs()

	benchmarks := []struct {
		name         string
		noOfElements int
	}{
		{"100", 100},
		{"1000", 1000},
		{"10000", 10000},
	}
	rand.Seed(time.Now().UnixNano())
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sortBenchmark(b, bm.noOfElements, sort)
			}
		})
	}
}

func sortBenchmark(b *testing.B, n int, sort func([]int) []int) []int {
	b.StopTimer()
	r := random(n)
	b.StartTimer()
	return sort(r)
}

func random(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = i
	}
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	return slice
}
