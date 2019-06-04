package main

func main() {
	cards := newDeckFromfile("my_cards")
	cards.shuffle()
	cards.print()
}
