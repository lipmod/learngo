package main


import (
	"fmt"
)

// new创建的过程是申请一块内存初始化给这个变量
func A() {
	p := new(int)
	fmt.Println(*p)
	fmt.Printf("type: %T\n", p)
	fmt.Printf("type: %T\n", *p)
	*p = 2
	fmt.Println(*p)
}

// 只是因为我看到这个*p = true有点奇怪
func B() {
	p := new(bool)
	*p = true
	fmt.Printf("type: %T,%T\n", p, *p)
}

// 命名返回值
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



func main(){
	A()
	B()
	C()
}