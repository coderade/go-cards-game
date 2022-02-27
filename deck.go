package main

import "fmt"

type deck []string

// receiver example in go, can be called as deck.print()
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
