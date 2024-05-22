package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const solutionLength = 5

// Game holds all the information we need to play a game of gordle.
type Game struct {
	reader *bufio.Reader
}

// New returns a Game, which can be used to Play.
func New(playerInput io.Reader) *Game {
	g := &Game{
		reader: bufio.NewReader(playerInput),
	}

	return g
}

// Play runs the game.
func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	guess := g.ask()

	fmt.Printf("Your guess is: %s\n", string(guess))
}

// ask reads input unitl a valid suggestion is made and returned.
func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", solutionLength)

	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %s\n", err.Error())
			continue
		}

		guess := splitToUpperCharacters(string(playerInput))

		err = g.validateGuess(guess)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution: %s.\n", err.Error())
		} else {
			return guess
		}
	}
}

// errInvalidWordLength is returned when the guess has the wrong number of characters.
var errInvalidWordLength = fmt.Errorf("invalid guess, word doesn't have teh same number of characters as the solution")

// validateGuess ensures the guess is valid enough.
func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != solutionLength {
		return fmt.Errorf("expected %d, got %d, %w", solutionLength, len(guess), errInvalidWordLength)
	}

	return nil
}

// splitToUpperCharacters is a naive implementation to turn a string into a list of characters
func splitToUpperCharacters(input string) []rune {
	return []rune(strings.ToUpper(input))
}
