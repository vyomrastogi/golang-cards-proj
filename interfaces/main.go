package main

import "fmt"

type bot interface {
	getGreeting() string
}

// TIL : if any type declares getGreeting and returns string,
// that type is member of type bot
// and the types can access printGreeting

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

//If reciever variable is not used, it can be omitted while declaring function
func (englishBot) getGreeting() string {
	//stimulate a custom logic for greeting
	return "Hello!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
