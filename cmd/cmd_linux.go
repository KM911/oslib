// go:build linux

package oslib

func Run(__COMAND__ string) {
	cmd = exec.Command("sh", "-c", __COMAND__)
	cmd.Run()
}

func RunStd(__COMAND__ string) {
	cmd = exec.Command("sh", "-c", __COMAND__)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func RunReturn(__COMAND__ string) string {
	cmd = exec.Command("sh", "-c", __COMAND__)
	out, _ := cmd.Output()
	return string(out)
}
