package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadGuess() (guess string, err error) {
	valid := false

	for !valid {
		fmt.Print("What is your letter? ")
		guess, err = reader.ReadString('\n')

		if err != nil {
			return guess, err
		}

		guess = strings.TrimSpace(guess)

		if len(guess) != 1 {
			fmt.Println("Please guess a single letter, no spaces.")
			continue
		}

		valid = true

	}
	return guess, nil
}
