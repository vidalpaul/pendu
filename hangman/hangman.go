package hangman

import "strings"

type Game struct {
	State        string   // "won", "lost", "playing"
	Letters      []string // letters in the word
	FoundLetters []string // letters found by the player
	UsedLetters  []string // letters used by the player
	TurnsLeft    int      // number of turns left
}

func New(turns int, word string) *Game {
	return &Game{
		State:        "playing",
		Letters:      strings.Split(strings.ToUpper(word), ""),
		FoundLetters: initFoundLetters(len(word), "_"),
		UsedLetters:  []string{},
		TurnsLeft:    turns,
	}
}

func initFoundLetters(n int, s string) []string {
	foundLetters := make([]string, n)
	for i := range foundLetters {
		foundLetters[i] = s
	}
	return foundLetters
}
