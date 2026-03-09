package main

import (
	"fmt"
)

// range 输出右边这个数组/切片的索引和值
func main() {
	notes := [7]int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println("len(notes)=", len(notes))

	for i := 0; i < len(notes); i++ {
		v := notes[i]
		fmt.Print(v)
	}
	fmt.Println()

	for i := range len(notes) {
		fmt.Print(notes[i])
	}
	fmt.Println()

	for i, note := range notes {
		fmt.Printf("%d%d ",i,note)
	}
	fmt.Println()

	for _, note1 := range notes {
		fmt.Print(note1)
	}
	fmt.Println()

	for range len(notes) {
		fmt.Print(1)
	}
	// go里面要求所有返回值都要被接收，但是又要求所有变量都要被使用，
	// 所以有了_ 用来占位并显式丢弃一个值
}

// range除了可以遍历数组，切片还能遍历map
