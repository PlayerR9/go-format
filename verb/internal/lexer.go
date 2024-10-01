package internal

import (
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/PlayerR9/go-errors/assert"
)

// Lexer is a lexer for format strings.
type Lexer struct {
	// stream is the stream to lex.
	stream io.RuneScanner

	// tokens is the list of tokens.
	tokens []*Token

	// prefix is the prefix of the format string.
	prefix rune

	// allowed_verbs is the list of allowed verbs.
	allowed_verbs []rune
}

// SetInputStream sets the stream to lex. Does nothing if the receiver
// is nil.
//
// Parameters:
//   - stream: The stream to lex.
func (l *Lexer) SetInputStream(stream io.RuneScanner) {
	if l == nil {
		return
	}

	l.stream = stream
}

// Lex lexes the stream. Does nothing if the receiver or the
// stream are nil.
//
// Returns:
//   - error: An error if the stream could not be lexed.
//
// NOTES:
//   - Remember to call Reset() after the stream has been lexed to avoid
//     previous tokens being left in the stream.
func (l *Lexer) Lex() error {
	if l == nil || l.stream == nil {
		return nil
	}

	for {
		tk, err := l.lex_one()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		if tk != nil {
			l.tokens = append(l.tokens, tk)
		}
	}

	return nil
}

// Tokens returns the list of tokens.
//
// Returns:
//   - []*Token: The list of tokens.
func (l Lexer) Tokens() []*Token {
	if len(l.tokens) == 0 {
		return nil
	}

	tokens := make([]*Token, 0, len(l.tokens))
	for _, tk := range l.tokens {
		assert.NotNil(tk, "tk must not be nil")

		tokens = append(tokens, tk)
	}

	return l.tokens
}

// Reset resets the lexer to make it reusable.
func (l *Lexer) Reset() {
	if l == nil {
		return
	}

	if len(l.tokens) > 0 {
		for i := 0; i < len(l.tokens); i++ {
			l.tokens[i] = nil
		}

		l.tokens = l.tokens[:0]
	}

	l.stream = nil
}

func NewLexer(prefix rune, allowed_verbs []rune) *Lexer {
	return &Lexer{
		prefix:        prefix,
		allowed_verbs: allowed_verbs,
	}
}

func (l *Lexer) lex_one() (*Token, error) {
	c, _, err := l.stream.ReadRune()
	if err != nil {
		return nil, err
	}

	var tk *Token

	if c == l.prefix {
		next_c, _, err := l.stream.ReadRune()
		if err == io.EOF && next_c == l.prefix {
			tk := NewToken(false, string(l.prefix))

			return tk, nil
		} else if err != nil {
			return nil, err
		}

		_, ok := slices.BinarySearch(l.allowed_verbs, next_c)
		if !ok {
			return nil, fmt.Errorf("flag \"%c%c\" is not supported", l.prefix, next_c)
		}

		tk = NewToken(true, string(next_c))
	} else {
		var builder strings.Builder

		builder.WriteRune(c)

		for {
			c, _, err := l.stream.ReadRune()
			if err == io.EOF {
				break
			} else if err != nil {
				return nil, err
			}

			if c == l.prefix {
				err := l.stream.UnreadRune()
				assert.Err(err, "l.stream.UnreadRune()")

				break
			}

			builder.WriteRune(c)
		}

		tk = NewToken(false, builder.String())
	}

	return tk, nil
}
