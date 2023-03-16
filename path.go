package lib

import (
	"os"
	"path/filepath"
)

/*
将文件路径转换为go文件路径
*/
func ToGoFile(path string) string {
	return path + ".go"
}

/*
返回文件的文件名 例如：/home/xxx/xxx.go 返回 /home/xxx/xxx
*/
func BaseName(__FILE__ string) string {
	return __FILE__[:len(__FILE__)-3]
}

/*
判断文件是否存在
*/
func IsExit(__FILE__ string) bool {

	_, err := os.Stat(__FILE__)
	if err != nil {
		println("文件路径为："+__FILE__, "文件不存在")
		return false
	}
	return true
}

/*
返回文件的绝对路径
*/
func AbsPath(__FILE__ string) string {
	abs, _ := filepath.Abs(__FILE__)
	return filepath.ToSlash(abs)
}

/*
获取当前命令窗口的路径
*/
func CmdPath() string {
	dir, _ := os.Getwd()
	return dir
}
