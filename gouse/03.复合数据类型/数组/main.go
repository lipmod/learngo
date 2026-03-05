package main

import (
	"fmt"
)

func main() {
	type Currency int
	const (
		USD Currency = iota
		EUR
	)
	symbol := [...]string{EUR: "$", USD: "€"}
	fmt.Println(EUR, symbol[EUR])
}

// USB EUR为字面量，当前的它们只是代表背后的数组，
// 然后这个symbol创建时用到的这个...其实就是让编译器推断这个数组最小要多大

// a := [...]string{
//     1: "B",
//     0: "A",
//     3: "D",
// }
// 长度 = 最大下标+1 = 4
// => ["A", "B", "", "D"]
