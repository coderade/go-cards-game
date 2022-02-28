package main

func main() {
	cards := newDeck()
	//hand, remainingDeck := deal(cards, 5)
	//hand.print()
	//fmt.Println("----")
	//remainingDeck.print()

	cards.saveToFile("my_card.txt")
}

//array => fixed length list of thing
//slice => list with length unfixed

// convert string to byte example
//greeting := "Hi there"
//byteString := []byte(greeting)
