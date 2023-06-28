package test

import (
	"github.com/KM911/oslib/fs"
	"testing"
)

func Test_Dir(t *testing.T) {

	t.Run("adt.InArray", func(t *testing.T) {
		if !fs.adt.InArray("a", []string{"a", "b", "c"}) {
			t.Error("adt.InArray")
		}
		if fs.adt.InArray("d", []string{"a", "b", "c"}) {
			t.Error("adt.InArray")
		}
	})

	t.Run("IsSameArray", func(t *testing.T) {
		list1 := []string{"a", "b", "c"}
		list2 := []string{"c", "b", "a"}
		list3 := []string{"a"}
		list4 := []string{"a", "b", "a"}
		if !fs.IsSameArray(list1, list2) || fs.IsSameArray(list1, list3) || fs.IsSameArray(list1, list4) {
			t.Error("IsSameArray")
		}
	})

	t.Run("test Dir", func(t *testing.T) {
		dirlist := fs.Dir(fs.CmdPath())
		mylist := []string{
			"dir_test.go",
			"cmd_test.go",
			"path_test.go",
		}
		if !fs.IsSameArray(dirlist, mylist) {
			t.Error("Dir")
		}
	})

	t.Run("test Dir", func(t *testing.T) {
		case1 := fs.FileName("C:\\Users\\Administrator\\Desktop\\test\\test.go")
		case2 := fs.FileName("test\\test.png")
		case3 := fs.FileName("test.txt")
		case4 := fs.FileName("test.asdfasdfasdfasdf")
		case5 := fs.FileName("asdfkdfjlkasdfjlasdfj/test")
		if case1 != case2 || case2 != case3 || case3 != case4 || case4 != case5 {
			t.Error("FileName")
		}
	})

	// 理论上来说 我们需要100%的代码覆盖率 还需要进行一个性能的分析就是说 还是不错的
}
