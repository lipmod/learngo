package main

import (
	"fmt"
)

func main() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "xi"}
	// 输出数组里面的 for循环1
	fmt.Println("len(notes)=", len(notes))
	for i := 0; i < len(notes); i++ {
		v := notes[i]
		fmt.Print(i, v, " ")
	}
	fmt.Println()

	// range 输出右边这个数组/切片的索引和值
	// 这里面判断终止的条件是len(notes)
	/*	for _,v := range s{
			...
		}
	// 本质上上面与下面是同一种写法	
		for i := 0; i < len(s); i++ {
	    	v := s[i]
	    	...
		}
	*/
	
	for index, note := range notes {
		fmt.Print(index, note, " ")
	}
	fmt.Println()

	for _, note1 := range notes {
		fmt.Print(note1, " ")
	}
	// go里面要求所有返回值都要被接收，但是又要求所有变量都要被使用，所以有了_，接收但直接抛弃
}

// range除了可以遍历数组，切片还能遍历map