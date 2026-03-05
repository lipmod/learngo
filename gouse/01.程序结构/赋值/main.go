package main

import "fmt"

func C2() (r int) {
	r = 10
	return
	// 不需要写返回值或者变量都行
}

func C1() (r int) {
	r = 10
	return 12
}

func C() {
	a := C2()
	b := C1()
	fmt.Println(a, b)
}

// 头一次看到这个赋值可以直接这样，之前看见这个多重赋值以为只有函数返回值用
// 当然这个在函数返回值主要是返回那个比尔值
func D() {
	a := 10
	b := 12
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
}

func main(){
	C()
	D()
}