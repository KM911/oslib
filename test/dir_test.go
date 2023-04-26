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

	t.Run("IsSameArray", func(t *testing.T) {
		list1 := []string{"a", "b", "c"}
		list2 := []string{"c", "b", "a"}
		list3 := []string{"a"}
		list4 := []string{"a", "b", "a"}
		if !oslib.IsSameArray(list1, list2) || oslib.IsSameArray(list1, list3) || oslib.IsSameArray(list1, list4) {
			t.Error("IsSameArray")
		}
	})

	t.Run("test Dir", func(t *testing.T) {
		dirlist := oslib.Dir(oslib.CmdPath())
		mylist := []string{
			"dir_test.go",
			"cmd_test.go",
			"path_test.go",
		}
		if !oslib.IsSameArray(dirlist, mylist) {
			t.Error("Dir")
		}
	})

	t.Run("test Dir", func(t *testing.T) {
		case1 := oslib.FileName("C:\\Users\\Administrator\\Desktop\\test\\test.go")
		case2 := oslib.FileName("test\\test.png")
		case3 := oslib.FileName("test.txt")
		case4 := oslib.FileName("test.asdfasdfasdfasdf")
		case5 := oslib.FileName("asdfkdfjlkasdfjlasdfj/test")
		if case1 != case2 || case2 != case3 || case3 != case4 || case4 != case5 {
			t.Error("FileName")
		}
	})

	// 理论上来说 我们需要100%的代码覆盖率 还需要进行一个性能的分析就是说 还是不错的
}
