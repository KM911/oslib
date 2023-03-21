package oslib

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
)

// 为了提升兼容性 我们使用的路径都是就是linux风格的
// 我们的文件换行都应该使用就是 LF
// 函数命名采用 驼峰命名法 首字母大写
// 变量命名采用 蛇形命名法 首字母小写
// 为了避免和系统包冲突 我们的变量名结尾都加上下划线 例如 file_ path_ type_ 等等 尽量不影响阅读

/*
将文件路类型换为特定类型的文件类型 返回相对路径 需要带 . 后缀
例如 /home/xxx/xxx.go 返回 /home/xxx/xxx.txt
test.mp4 返回 test.mp3
*/
func TransformFileType(path string, type_ string) string {
	if type_[:1] != "." {
		type_ = "." + type_
	}
	if filepath.Ext(path) == type_ {
		return path
	} else {
		return path[:len(path)-len(filepath.Ext(path))] + type_
	}
}

/*
返回文件的文件名

例如：/home/xxx/xxx.go 返回 /home/xxx/xxx

test.go 返回 /home/xxx/test

对于path.Dir()方法 区别就是可以对于相对路径也返回它的绝对路径的文件夹路径
需要文件名包括后缀 应该使用filepath.Base(file_)方法
*/
func DirName(file_ string) string {
	if filepath.IsAbs(file_) {
		return file_[:len(file_)-len(filepath.Ext(file_))]
	} else {
		abs_file_, _ := filepath.Abs(file_)
		return abs_file_[:len(abs_file_)-len(filepath.Ext(abs_file_))]
	}
}

/*
返回文件名 不包含后缀
例如 /home/bin/xxx.go 返回 xxx
test.go 返回 test
*/
func FileName(file_ string) string {
	return filepath.Base(file_)[:len(filepath.Ext(file_))]
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

/*
返回文件的绝对路径
移除错误处理
*/
func AbsPath(file_ string) string {
	abs, _ := filepath.Abs(file_)
	return filepath.ToSlash(abs)
}

/*
获取当前命令窗口的路径
移除错误处理
*/
func CmdPath() string {
	dir, _ := os.Getwd()
	return filepath.ToSlash(dir)
}

/*
获取运行时的路径 这里会和cmdpath不同吗?
*/
func RunPath() string {
	_, fullFilename, _, _ := runtime.Caller(0)
	return path.Dir(fullFilename)
}

/*
判断该字符串是否在数组中
*/
func InArray(string_ string, list []string) bool {
	if len(list) == 0 {
		return false
	}
	for _, i := range list {
		if i == string_ {
			return true
		}
	}
	return false
}
