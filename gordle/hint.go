package gordle

// hint describes the validity of a character in a word
type hint byte

const (
	absentCharacter hint = iota
	wrongPosition
	correctPosition
)
