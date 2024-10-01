package internal

// Token is a token in a format string.
type Token struct {
	// IsVerb specifies if the token is a verb.
	IsVerb bool

	// Data is the data of the token.
	Data string
}

// NewToken creates a new token.
//
// Parameters:
//   - is_verb: Specifies if the token is a verb.
//   - data: The data of the token.
//
// Returns:
//   - *Token: The new token. Never returns nil.
func NewToken(is_verb bool, data string) *Token {
	return &Token{
		IsVerb: is_verb,
		Data:   data,
	}
}
