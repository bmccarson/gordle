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

// feedback is a list of hints, on per character of the word
type feedback []hint

// StringConcat is a naive implementation to build feedback as a string.
// It is used only to benchmark it against the string.Builder version.
func (fb feedback) StringConcat() string {
	var output string
	for _, h := range fb {
		output += h.String()
	}
	return output
}
