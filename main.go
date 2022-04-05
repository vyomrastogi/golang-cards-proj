package main

import "fmt"

func main() {

	//card := "Ace of Spades" // This initializes card as string and assigns value
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades") //appends returns a new slice and doesn't change value of existing cards

	cards := newDeck()
	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()

	//test toString o/p
	fmt.Print(cards.toString())

	//save original deck to file
	cards.saveToFile("my_cards")

	//load new deck from file
	cardsFromFile := newDeckFromFile("my_cards")
	cardsFromFile.print()

	//load a non-existent file to exit program
	//commented
	//newDeckFromFile("i_do_not_exist")

}

// func newCard() string {
// 	return "Five of Diamond"
// }
