package main // make this executable

import "fmt" // Importing code from another package

// This is a simple Go program that prints "Hello, World!" to the console.
// The main function serves as the entry point of the program.

func main() {
	// variables can be initialized inside a funtion only. It can be declared at package level but not initialized. Assigning value is the next step.
	// Variables in Go are statically typed, meaning their type is known at compile time.
	var card string = "Ace of Spades" // Declare a variable named card and assign it a value. card will hold a string value (static type like java, c++). In js, pytho, ruby it would be var card = "Ace of Spades" and card can hold string or other types of values (dynamic type).
	// card := "Ace of Spades" // This is also correct and we delegate the type to go compiler. Only when declaring new variable
	fmt.Println(card) // Print the value of card to the console
	card = newCard()
	fmt.Println("Hello, World!")

	// Arrays
	var cards []string = []string{newCard(), newCard()}
	// OR
	//cards := []string{"Ace of Spades", "Two of Hearts", "Three of Diamonds"} // Declare and initialize an array of strings
	fmt.Println(cards)                      // Print the array of cards to the console
	cards = append(cards, "King of Hearts") // Append doesn't modify the original array but returns a new one
	fmt.Println(cards)                      // Print the updated array of cards

	for i, c := range cards { // range to Iterate over the array of cards
		fmt.Printf("Card %d: %s\n", i+1, c) // Print each card with its index
	}
	fmt.Println("Total cards:", len(cards)) // Print the total number of cards

}
func newCard() string {
	return "Five of Diamonds"
}
