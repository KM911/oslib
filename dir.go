package oslib

import (
	"os"
	"path/filepath"
)

/*
返回文件夹下的所有文件 不包括子文件夹 和隐藏文件
*/
func Dir(__DirPath__ string) []string {
	files, err := os.ReadDir(__DirPath__)
	if err != nil {
		panic(err)
	}
	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	return fileNames
}

/*
返回文件夹下的所有文件 不包括特定后缀的文件
*/
func Dir_Ingore(__DirPath__ string, __Tyee__ []string) []string {
	files, err := os.ReadDir(__DirPath__)
	if err != nil {
		panic(err)
	}
	var fileNames []string
	for _, file := range files {
		if InArray(filepath.Ext(file.Name()), __Tyee__) {
			fileNames = append(fileNames, file.Name())
		}
	}
	return fileNames
}

/*
返回文件夹下的所有文件 只包括特定后缀的文件
*/
func Dir_Keep(__DirPath__ string, __Tyee__ []string) []string {
	files, err := os.ReadDir(__DirPath__)
	if err != nil {
		panic(err)
	}
	var fileNames []string
	for _, file := range files {
		if InArray(filepath.Ext(file.Name()), __Tyee__) {
			fileNames = append(fileNames, file.Name())
		}
	}
	return fileNames
}

/*
判定文件是否为目录
*/
func IsDir(__FILE__ string) bool {
	fileInfo, err := os.Stat(__FILE__)
	if err != nil {
		panic(err)
	}
	return fileInfo.IsDir()
}
