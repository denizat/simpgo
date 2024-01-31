package tting

import (
	// "fmt"
	"testing"
)


var global string

func BenchmarkStr(b * testing.B) {
	var res string
	richard := "richard"
	for n := 0; n < b.N; n++ {
		res = indexstr(richard)
	}
	global = res

}
func BenchmarkFr(b * testing.B) {
	var res []byte
	richard := "richard"
	for n := 0; n < b.N; n++ {
		res = indexfr(richard)
	}
	global = string(res)
}
