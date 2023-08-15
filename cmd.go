package oslib

import (
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

var (
	cmd *exec.Cmd
)

/*
无声执行指令 no blocking
*/
func Run(__COMAND__ string) {
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", __COMAND__)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	} else {
		cmd = exec.Command("sh", "-c", __COMAND__)
	}
	cmd.Run()
}

/*
执行指令 并重定向输出
*/
func RunStd(__COMAND__ string) {
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", __COMAND__)
	} else {
		cmd = exec.Command("sh", "-c", __COMAND__)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

/*
执行指令并返回结果 blocking
*/
func RunReturn(__COMAND__ string) string {
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", __COMAND__)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	} else {
		cmd = exec.Command("sh", "-c", __COMAND__)
	}
	out, _ := cmd.Output()
	return string(out)
}
