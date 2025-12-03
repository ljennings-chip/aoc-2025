package main

import (
	"testing"
)

func BenchmarkSolve(b *testing.B) {
	b.ReportAllocs()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total := 0
		for off := 0; off+100 <= len(data); off += 100 + 1 {
			total += TopKDigitsOrdered(data[off : off+100])
		}
	}

	// Add a custom microseconds-per-op metric
	usPerOp := float64(b.Elapsed().Nanoseconds()) / float64(b.N) / 1e3
	b.ReportMetric(usPerOp, "us/op")
}
