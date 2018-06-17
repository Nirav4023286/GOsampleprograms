package main

func main() {
	//var card string = "Ace of spades"  // var we are about to create a new variable
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	// how we can iterate
	cards.print()
}

func newCard() string { //newCard- define a function called 'newCard'
	// string-->  when executed this function will return a value of type 'string'
	return "Five of Diamonds1"
}
