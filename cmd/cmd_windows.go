package oslib

import (
	"os"
	"os/exec"
	"syscall"
)

func Run(__COMMAND__ string) {
	cmd = exec.Command("cmd", "/C", __COMMAND__)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Run()
}

func RunStd(__COMMAND__ string) {
	cmd = exec.Command("cmd", "/C", __COMMAND__)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func RunReturn(__COMMAND__ string) string {
	cmd = exec.Command("cmd", "/C", __COMMAND__)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, _ := cmd.Output()
	return string(out)
}
