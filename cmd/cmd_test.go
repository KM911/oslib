package oslib

import "testing"

// 这里我们开始测试函数传递的区别

// 内联函数 主要是会进入程序的栈中

func SimpleFunction(a, b int) int {
	return a + b
}

// 传递值和传递指针的区别
func SimpleFunction2(a, b *int) int {
	return *a + *b
}

func Swap(a, b int) (int, int) {
	return b, a
}
func Swap1(a, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return b, a
}

// 成功修改了值
func Swap2(a, b *int) {
	*a, *b = *b, *a
}

// 这里就是指针 你修改了后会影响到原来的值
func SwapObject(a []int) {
	a[0], a[1] = a[1], a[0]
}

func SwapObject1(a *[]int) {
	(*a)[0], (*a)[1] = (*a)[1], (*a)[0]
}

func TestFunctionPointer(t *testing.T) {
	a := 1
	b := 2
	Swap(a, b)
	println(a, b)
}

// C++里会拷贝对象呗 就是说
func TestFunctionPointer1(t *testing.T) {
	a := 1
	b := 2
	Swap1(a, b)
	println(a, b)
}

func TestFunctionPointer2(t *testing.T) {
	arr := []int{1, 2}
	SwapObject(arr)
	println(arr[0], arr[1])
}

func TestFunctionPointer3(t *testing.T) {
	arr := []int{1, 2}
	SwapObject1(&arr)
	println(arr[0], arr[1])
}

//func BenchmarkSimpleFunction(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		_ = SimpleFunction(1998, 265)
//	}
//}
//
//func BenchmarkSimpleFunction2(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		_ = 1998 + 265
//	}
//}
