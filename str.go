package oslib

// 像这样的一行代码是不是会有性能差别 我还挺好奇的就是说
func GetSubfix(str string, num int) string {
	return str[len(str)-num:]
}

func GetPrefix(str string, num int) string {
	return str[:num]
}
