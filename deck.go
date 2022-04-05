package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of deck
// which is a slice of strings

// the deck 'type' has all the properties of slice of string
type deck []string

//generates and returns a new deck of 15 cards
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"One", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

//by convention in a reciever method variable name is single character of it's type
// Reciever function: prints the cards in the deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//go can return multiple values
// deal the deck for given handSize
func deal(d deck, handSize int) (deck, deck) {
	//slice creates two references to new slices and
	//doesn't modify existing slice
	return d[:handSize], d[handSize:]
}

//save the deck to provided filename
//returns the error returned by ioutil.WriteFile
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//creates a new deck from file
//exits program if reading file results in an error
func newDeckFromFile(filename string) deck {
	bytesize, err := ioutil.ReadFile(filename)
	//check if error is nil or not
	if err != nil {
		//print the error and exit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	//convert the byte size to deck
	deckslice := strings.Split(string(bytesize), ",")
	return deck(deckslice)
}

//shuffles the existing deck
func (d deck) shuffle() {
	//randomize the seed source of 
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		//swap positions
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

//returns string representation of deck
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
