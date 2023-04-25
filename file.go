package oslib

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

/*
返回文件名 不包含后缀
例如 /home/bin/xxx.go 返回 xxx
test.go 返回 test
*/
func FileName(file_ string) string {
	return filepath.Base(file_)[:len(filepath.Ext(file_))]
}

// GetFileNameWithoutExtension 获取不包含扩展名的文件名
func GetFileNameWithoutExtension(filePath string) string {
	fileName := filepath.Base(filePath)
	dotIndex := strings.LastIndex(fileName, ".")
	if dotIndex == -1 {
		return fileName
	}
	return fileName[:dotIndex]
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

// GetFileName 获取文件名，不包括路径和扩展名
func GetFileName(filePath string) (fileName string) {
	absPath, _ := filepath.Abs(filePath)
	fileBase := filepath.Base(absPath)
	fileExt := filepath.Ext(fileBase)
	fileName = fileBase[:len(fileBase)-len(fileExt)]
	return fileName
}

/*
判断文件或者文件夹是否存在
*/
func IsExit(file_ string) bool {
	_, err := os.Stat(file_)
	if err != nil {
		return false
	}
	return true
}

func Check(a any) {
	fmt.Println("type is ", reflect.TypeOf(a))
	fmt.Println("data is ", a)
}

/*
将文件路类型换为特定类型的文件类型 返回相对路径 需要带 . 后缀
例如 /home/xxx/xxx.go 返回 /home/xxx/xxx.txt
test.mp4 返回 test.mp3
*/
func ChangeFileType(path string, type_ string) string {
	ext := filepath.Ext(path)
	withExt := strings.TrimSuffix(path, ext)
	return withExt + "." + type_
}
