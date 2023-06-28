package adt

/*
判断该字符串是否在数组中
*/
func InArray(string_ string, list []string) bool {
	if len(list) == 0 {
		return false
	}
	for _, i := range list {
		if i == string_ {
			return true
		}
	}
	return false
}

/*
判断两个字符串数组是否相等
长度相同
内容相同
不要求内容相同
*/
func IsSameArray(array1 []string, array2 []string) bool {
	if len(array1) != len(array2) {
		return false
	}
	for _, i := range array1 {
		if !InArray(i, array2) {
			return false
		}
	}
	return true
}
