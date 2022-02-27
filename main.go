package main

func main() {
	cards := deck{"Ace of diamonds"}
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}

//array => fixed length list of thing
//slice => list with length unfixed
