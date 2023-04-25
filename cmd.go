package oslib

import (
	"os"
	"os/exec"
	"runtime"
)

/*
无声执行指令 no blocking
*/
func Run(__COMAND__ string) {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/C", __COMAND__)
		cmd.Run()
	} else {
		cmd := exec.Command("sh", "-c", __COMAND__)
		cmd.Run()
	}
}

/*
执行指令 并重定向输出
*/

func RunStd(__COMAND__ string) {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/C", __COMAND__)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	} else {
		cmd := exec.Command("sh", "-c", __COMAND__)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
}

/*
执行指令并返回结果 blocking
*/
func RunReturn(__COMAND__ string) string {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/C", __COMAND__)
		out, _ := cmd.Output()
		return string(out)
	} else {
		cmd := exec.Command("sh", "-c", __COMAND__)
		out, _ := cmd.Output()
		return string(out)
	}
}

/*
可以捕获报错 提高我们的检索性
*/
func RunReturnError(command string) (string, error) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}
	out, err := cmd.Output()
	return string(out), err
}

