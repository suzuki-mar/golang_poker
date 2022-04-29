package main

import (
	"fmt"
	"golang/src/github.com/suzuki-mar/hand"
)

type Hander interface {
	String() string
}

func main() {
	hand := hand.NewHand()

	printHand(hand)
}

func printHand(hand Hander) {
	fmt.Println(hand.String())
}
