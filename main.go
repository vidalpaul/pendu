package main

import (
	"fmt"
	"os"
	"vidalpaul/pendu/hangman"
)

func main() {
	hangman.DrawWelcome()

	g := hangman.New(8, "Golang")

	fmt.Println(g)

	l, err := hangman.ReadGuess()

	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	fmt.Println("Letter guessed:", l)
}
