package oslib

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
返回文件的文件名
例如：/home/xxx/xxx.go 返回 /home/xxx/xxx
test.go 返回 /home/xxx/test
需要文件名包括后缀 应该使用filepath.Base(__FILE__)方法
*/
func DirName(__FILE__ string) string {
	if filepath.IsAbs(__FILE__) {
		return __FILE__[:len(__FILE__)-len(filepath.Ext(__FILE__))]
	} else {
		__ABSFILE__, _ := filepath.Abs(__FILE__)
		return __ABSFILE__[:len(__ABSFILE__)-len(filepath.Ext(__ABSFILE__))]
	}
}

/*
返回文件名 不包含后缀
例如 /home/xxx/xxx.go 返回 xxx
test.go 返回 test
*/
func FileName(__FILE__ string) string {
	return filepath.Base(__FILE__)[:len(filepath.Ext(__FILE__))]
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
移除错误处理
*/
func AbsPath(__FILE__ string) string {
	abs, _ := filepath.Abs(__FILE__)
	return filepath.ToSlash(abs)
}

/*
获取当前命令窗口的路径
移除错误处理
*/
func CmdPath() string {
	dir, _ := os.Getwd()
	return dir
}

/*
判断该字符串是否在数组中
*/
func InArray(__STRING__ string, __ARRAY__ []string) bool {
	if len(__ARRAY__) == 0 {
		return false
	}
	for _, i := range __ARRAY__ {
		if i == __STRING__ {
			return true
		}
	}
	return false
}
