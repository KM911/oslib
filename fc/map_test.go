package fc

import (
	"fmt"
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	stringSlice := []string{"A", "B", "C"}
	fmt.Println(Map(stringSlice, strings.ToLower))
}

func BenchmarkMap(b *testing.B) {
	stringSlice := []string{"A", "B", "C"}
	for i := 0; i < b.N; i++ {
		Map(stringSlice, strings.ToLower)
	}
}
func BenchmarkMap1(b *testing.B) {
	stringSlice := []string{"A", "B", "C"}
	for i := 0; i < b.N; i++ {
		Map1(stringSlice, strings.ToLower)
	}
}
func BenchmarkMapMethod(b *testing.B) {
	stringSlice := []string{"A", "B", "C"}
	for i := 0; i < b.N; i++ {
		newSlice := make([]string, len(stringSlice))
		for index, value := range stringSlice {
			newSlice[index] = strings.ToLower(value)
		}
	}
}
