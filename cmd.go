package lib

import (
	"os"
	"os/exec"
	"runtime"
)

/*
直接执行指令
*/
func Run(__COMAND__ string) {
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
执行指令并返回结果
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
