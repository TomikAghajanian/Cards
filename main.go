package main

import(
	"fmt"
)

func main() {
	cards := newDeckFromFile("my_card")
	cards.shuffle()
	fmt.Println(cards.toString())
}