package oslib

import (
	"os"
	"path/filepath"
)

/*
返回文件夹下的所有文件 不包括子文件夹 和隐藏文件
*/
func Dir(path string) []string {

	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	var fileNames []string
	for _, file := range files {
		// 我们不希望就是显示文件夹 而是只显示文件
		if file.IsDir() {
			continue
		}
		fileNames = append(fileNames, file.Name())

	}
	return fileNames
}

/*
返回文件夹下的所有文件 包括就是子文件夹下的文件
例如 [a.txt b.txt c.txt a/a.txt a/b.txt a/c.txt]
*/
func DeepDir(path string) []string {
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	var fileNames []string
	for _, file := range files {
		if file.IsDir() {
			deepFileNames := DeepDir(path + "/" + file.Name())
			for _, deepFileName := range deepFileNames {
				fileNames = append(fileNames, file.Name()+"/"+deepFileName)
			}

		} else {
			fileNames = append(fileNames, file.Name())
		}
	}
	return fileNames
}

/*
返回文件夹下的所有文件 不包括特定后缀的文件
*/
func DirIngore(path string, type_ []string) []string {
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	var fileNames []string
	for _, file := range files {
		if InArray(filepath.Ext(file.Name()), type_) {
			fileNames = append(fileNames, file.Name())
		}
	}
	return fileNames
}

/*
返回文件夹下的所有文件 只包括特定后缀的文件
*/
func DirKeep(path string, type_ []string) []string {
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	var fileNames []string
	for _, file := range files {
		if InArray(filepath.Ext(file.Name()), type_) {
			fileNames = append(fileNames, file.Name())
		}
	}
	return fileNames
}

/*
判定文件是否为目录
*/
func IsDir(file_ string) bool {
	fileInfo, err := os.Stat(file_)
	if err != nil {
		panic(err)
	}
	return fileInfo.IsDir()
}
