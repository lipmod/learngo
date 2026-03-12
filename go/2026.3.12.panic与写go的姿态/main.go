package main

import (
	"fmt"
	"math/rand"
)

type suit string

const (
	Spades   suit = "Spades"
	Hearts   suit = "Hearts"
	Diamonds suit = "Diamonds"
	Clubs    suit = "Clubs"
	Joker    suit = "joker"
)

func drawCard() suit {
	suits := []suit{Spades, Hearts, Diamonds, Clubs, Joker}
	return suits[rand.Intn(len(suits))]
}

func main() {
	switch s := drawCard(); s {
	case Spades:
		fmt.Println("黑桃")
	case Hearts:
		fmt.Println("红桃")
	case Diamonds:
		fmt.Println("方块")
	case Clubs:
		fmt.Println("梅花")
	default:
		panic(fmt.Sprintf("invalid suit %q", s))
	}
}

// 这个程序其实没什么意义，只是在这个程序才理解到panic是发生本不应该发生的错误的处理方式
