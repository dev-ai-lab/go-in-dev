package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()     // Call the print method to display the cards
	println(cards[0]) // Print the first card directly

	hand, remainingDeck := deal(cards, len(cards)-2) // Deal 5 cards from the deck
	hand.print()
	remainingDeck.print() // Print the remaining deck after dealing

	println(cards.toString())
	cards.saveToFile("my_cards.txt")                // Save the deck to a file
	deckFromFile := newDeckFromFile("my_cards.txt") // Load the deck from a file (not implemented
	fmt.Print(deckFromFile.toString())              // Print the loaded deck

	ch := make(chan string) // create a channel

	go sayHi(ch) // start goroutine

	msg := <-ch // receive message
	fmt.Println(msg)

}

func sayHi(ch chan string) {
	// This function runs in a separate goroutine. Time consuming task running in the background
	ch <- "Hi from goroutine!" // send message to channel
}
