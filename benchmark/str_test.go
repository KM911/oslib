package benchmark

import (
	"github.com/KM911/oslib"
	"testing"
)

func BenchmarkStrByFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		println(oslib.GetPrefix("dzg", 2))
		println(oslib.GetPrefix("commondzg", 2))
	}
}

func BenchmarkStrByCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		println("dzg"[:2])
		println("commondzg"[:2])
	}
}
