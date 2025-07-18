package main

import (
	"fmt"
	"os"
	"strings"
)

type Deck []string // Define a new type Deck which extends a slice of strings or borrows from it.

func newDeck() Deck { // Function to create a new deck of cards
	cards := Deck{}
	cards = append(cards, "Ace of Spades", "Two of Hearts", "Three of Diamonds", "Two of Spades") // Initialize the deck with some cards
	return cards                                                                                  // Return the deck of cards
}
func (d Deck) print() { // Method to print the deck of cards. A function with a receiver. It means: any variable of type Deck can access this method. or in OOPS world, card instance passed to the fuciton as a reference (this, self) or think any Card is instance had a method called print. As go convetion, use single letter for receiver variable.
	for i, card := range d { // if i is not used, use _ instead of i.
		fmt.Printf("Card %d: %s\n", i+1, card) // Print each card with its index
	}
}

func deal(d Deck, handSize int) (Deck, Deck) { // Function to deal cards from the deck
	return d[:handSize], d[handSize:] // Return two slices: one for the hand and one for the remaining deck
}

// some io functions

func (d Deck) saveToFile(filename string) error {
	err := os.WriteFile(filename, []byte(d.toString()), 0666) // Save the deck to a file
	return err
}
func newDeckFromFile(filename string) Deck {
	byteSlice, err := os.ReadFile(filename)
	if err != nil {
		fmt.Print("Error reading file: ", err)
		os.Exit(1) // Exit the program if there is an error
	}
	cards := strings.Split(string(byteSlice), ",") // Convert the byte slice to a string and split it into cards
	return Deck(cards)                             // Return the deck of cards
}

func (d Deck) toString() string {
	return strings.Join([]string(d), ",") // Convert the deck to a string representation
}
