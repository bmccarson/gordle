package gordle

// Game holds all the information we need to play a game of gordle.
type Game struct{}

// New returns a Game, which can be used to Play.
func New() *Game {
	g := &Game{}

	return g
}
