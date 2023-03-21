package test

import (
	"fmt"
	"github.com/KM911/oslib"
	"testing"
)

func Test_cmd(t *testing.T) {

}

func Test_Dir(t *testing.T) {

	t.Run("InArray", func(t *testing.T) {
		if !oslib.InArray("a", []string{"a", "b", "c"}) {
			t.Error("InArray")
		}
		if oslib.InArray("d", []string{"a", "b", "c"}) {
			t.Error("InArray")
		}

	})
	t.Run("Dir", func(t *testing.T) {
		fmt.Println(oslib.Dir("D:\\SOFT\\BIN"))

	})

	t.Run("cmdpath", func(t *testing.T) {
		fmt.Println(oslib.CmdPath())
		fmt.Println(oslib.RunPath())
	})
}
