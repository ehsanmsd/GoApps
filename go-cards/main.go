package main

// func main() {
// 	cards := newDeck()
// 	hand, remainingDeck := deal(cards, 5)
// 	hand.print()
// 	remainingDeck.print()
// }

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
