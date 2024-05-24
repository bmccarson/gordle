package gordle

// hint describes the validity of a character in a word
type hint byte

const (
	absentCharacter hint = iota
	wrongPosition
	correctPosition
)

func (h hint) String() string {
	switch h {
	case absentCharacter:
		return "◻️" // grey square
	case wrongPosition:
		return "🟡" // yellow circle
	case correctPosition:
		return "💚" // green heart
	default:
		// shouldn't happen, if so its an error
		return "💔" // red broken heart
	}
}
