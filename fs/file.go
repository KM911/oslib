package fs

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

/*
返回文件名 不包含后缀 对于需要文件名包括后缀 应该使用filepath.Base(file_)方法
例如 /home/bin/xxx.go 返回 xxx
test.go 返回 test
*/
func FileName(file_ string) string {
	ext := filepath.Ext(file_)
	file_ = filepath.Base(file_)
	length := len(file_)
	return file_[:length-len(ext)]
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
判断文件或者文件夹是否存在
*/
func IsExit(file_ string) bool {
	_, err := os.Stat(file_)
	if err != nil {
		return false
	}
	return true
}

/*
将文件类型换为特定类型的文件类型 返回相对路径 不需要带 . 后缀
例如 ChangeFileType("/home/xxx/xxx.go","txt") 返回 /home/xxx/xxx.txt
ChangeFileType("xxx","mp4") 返回 xxx.mp4
*/
func ChangeFileType(path string, type_ string) string {
	ext := filepath.Ext(path)
	withExt := strings.TrimSuffix(path, ext)
	return withExt + "." + type_
}

/*
检查文件类型是否为特定类型
例如 main.go go true
main.go txt false
*/
func IsFileType(file string, type_ string) bool {
	return filepath.Ext(file) == "."+type_
}

/*
读取文本文件
*/
func Read(file string) string {
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
func Create(src string, data string) {
	file, _ := os.Create(src)
	defer file.Close()
	file.WriteString(data)
}

func Append(src string, data string) {
	file, _ := os.OpenFile(src, os.O_RDWR|os.O_APPEND, 0644)
	defer file.Close()
	file.WriteString(data)
}

// TODO use the panic or error ????
func Move(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("Couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("Writing to output file failed: %s", err)
	}
	// The copy was successful, so now delete the original file
	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("Failed removing original file: %s", err)
	}
	return nil
}

func Copy(src, dst string) error {
	inputFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(dst)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("Couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("Writing to output file failed: %s", err)
	}
	return nil
}

func Remove(src string) {
	syscall.Unlink(src)
	syscall.Rmdir(src)
}
