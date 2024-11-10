package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()

	number := []int{3, 2, 1}
	getAvg(number)
}
