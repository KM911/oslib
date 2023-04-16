package oslib

import (
	"os"
	"path/filepath"
)

/*
返回文件名 不包含后缀
例如 /home/bin/xxx.go 返回 xxx
test.go 返回 test
*/
func FileName(file_ string) string {
	return filepath.Base(file_)[:len(filepath.Ext(file_))]
}

/*
返回文件的文件名 携带前缀的路径 并且一定为绝对路径

例如：/home/xxx/xxx.go 返回 /home/xxx/xxx

test.go 返回 /home/xxx/test

对于path.Dir()方法 区别就是可以对于相对路径也返回它的绝对路径的文件夹路径
需要文件名包括后缀 应该使用filepath.Base(file_)方法
*/
func FullFileName(file_ string) string {
	if filepath.IsAbs(file_) {
		return file_[:len(file_)-len(filepath.Ext(file_))]
	} else {
		abs_file_, _ := filepath.Abs(file_)
		return abs_file_[:len(abs_file_)-len(filepath.Ext(abs_file_))]
	}
}

/*
判断文件是否存在
*/
func IsExit(file_ string) bool {
	_, err := os.Stat(file_)
	if err != nil {
		println("文件路径为："+file_, "文件不存在")
		return false
	}
	return true
}
