package main

import (
	"os"

	"github.com/bmccarson/gordle/gordle"
)

const maxAttempts = 6

func main() {
	solution := "hello"

	g := gordle.New(os.Stdin, solution, maxAttempts)

	g.Play()
}
