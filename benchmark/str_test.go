package benchmark

import (
	"github.com/KM911/oslib/adt"
	"testing"
)

func BenchmarkStrByFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		println(adt.GetPrefix("dzg", 2))
		println(adt.GetPrefix("commondzg", 2))
	}
}

func BenchmarkStrByCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		println("dzg"[:2])
		println("commondzg"[:2])
	}
}
