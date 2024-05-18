package main

import (
	"os"

	"github.com/bmccarson/gordle/gordle"
)

func main() {
	g := gordle.New(os.Stdin)
	g.Play()
}
