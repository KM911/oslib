package oslib

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
)

/*
返回文件名 不包含后缀
例如 /home/bin/xxx.go 返回 xxx
test.go 返回 test
*/
func FileName(file_ string) string {
	ext := filepath.Ext(file_)
	file_ = filepath.Base(file_)
	length := len(file_)
	return file_[:length-len(ext)]
	// 其实这里我还是像做一个对比的 就是看看我们的编译器是不是对其一些优化
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

/*

 */
// 从现在开始不是使用return 而是使用有意义的变量名?
func RandomNumberName(length int) string {
	return strconv.FormatInt(time.Now().Unix(), length)
}

var StringTable = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"a", "b", "c", "d", "e", "f", "g", "h", "i",
	"j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

/*
base64是有 10 + 26 +26 = 62个字符组成的
*/
func RandomStringName(length int) (StringName string) {
	rand.Seed(time.Now().UnixNano())
	StringBuilder := strings.Builder{}
	for i := 0; i < length; i++ {
		StringBuilder.WriteString(StringTable[rand.Intn(62)])
	}
	StringName = StringBuilder.String()
	return StringName
}

/*
读取文本文件
*/
func ReadTextFile(file string) string {
	open, err := os.Open(file)
	if err != nil {
		return ""
	}
	defer open.Close()
	stat, err := open.Stat()
	if err != nil {
		return ""
	}
	size := stat.Size()
	var data = make([]byte, size)
	open.Read(data)
	return string(data)
}

/*
检查文件类型是否为特定类型
例如 main.go go true
main.go txt false
*/
func IsFileType(file string, type_ string) bool {
	return filepath.Ext(file) == "."+type_
}

// python 真的为我做了很多 就是说 在linux中是使用
// 这里应该返回一个迭代器的就是说 我们这里还需要就是
//func ReadTextByline(file string) string {
//
//}
