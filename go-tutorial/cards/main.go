package main

func main() {
	cards := newDeck()
	cards.shuffle()
	hand1, remaining := deal(cards, 5)
	hand2, _ := deal(remaining, 5)
	hand1.print()
	hand2.print()
}
