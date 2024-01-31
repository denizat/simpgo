package tting

import (
	"testing"
)

func BenchmarkStr(b * testing.B) {
	for n := 0; n < b.N; n++ {
		indexstr("richard")
	}
}
func BenchmarkFr(b * testing.B) {
	for n := 0; n < b.N; n++ {
		indexfr("richard")
	}
}
