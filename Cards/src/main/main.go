package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	hand, _ := deal(cards, 5)
	hand.saveToFile("hand5")
	newHand := newDeckFromFile("hand5")
	newHand.print()
}
