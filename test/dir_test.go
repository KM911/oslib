package test

import (
	"github.com/KM911/oslib"
	"testing"
)

func Test_Dir(t *testing.T) {

	t.Run("InArray", func(t *testing.T) {
		if !oslib.InArray("a", []string{"a", "b", "c"}) {
			t.Error("InArray")
		}
		if oslib.InArray("d", []string{"a", "b", "c"}) {
			t.Error("InArray")
		}

	})

}
