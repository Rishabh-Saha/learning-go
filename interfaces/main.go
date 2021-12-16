package main

import "fmt"

type bot interface {
	getGreetring() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreetring())
}

func (englishBot) getGreeting() string {
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
