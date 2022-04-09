package main

import (
	"fmt"
	"golang/src/github.com/suzuki-mar/poker"
)

type Hander interface {
	String() string
}

func main() {

	hand := poker.BuildHand()

	printHand(hand)
}

func printHand(hand Hander) {
	fmt.Println(hand.String())
}
