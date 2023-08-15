package fc

// 利用反射

func Map[T, R any](array []T, callback func(T) R) []R {
	result := make([]R, len(array))
	for i := 0; i < len(array); i++ {
		result[i] = callback(array[i])
	}
	return result
}

// 这里再次说明了我们不应该使用append 这是一个很耗时的操作
func Map1[T, R any](array []T, callback func(T) R) []R {
	var result []R
	for _, value := range array {
		result = append(result, callback(value))
	}
	return result
}
