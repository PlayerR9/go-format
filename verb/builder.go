package verb

import (
	"bytes"
	"slices"

	"github.com/PlayerR9/go-format/verb/internal"
)

const (
	// DefaultPrefix is the default prefix for format strings.
	DefaultPrefix rune = '%'
)

// Builder is a builder for format strings.
type Builder struct {
	// prefix is the prefix for specifying format verbs.
	prefix rune

	// allowed_verbs is the list of allowed format verbs.
	allowed_verbs []rune
}

// SetPrefix sets the prefix for specifying format verbs. Does
// nothing if the receiver is nil.
//
// Parameters:
//   - char: The prefix to set.
func (b *Builder) SetPrefix(char rune) {
	if b == nil {
		return
	}

	b.prefix = char
}

// Register registers a format verb. Does nothing if the receiver
// is nil.
//
// Verbs that are the same as the prefix are ignored.
//
// Parameters:
//   - verb: The format verb to register.
func (b *Builder) Register(verb rune) {
	if b == nil {
		return
	}

	pos, ok := slices.BinarySearch(b.allowed_verbs, verb)
	if !ok {
		b.allowed_verbs = slices.Insert(b.allowed_verbs, pos, verb)
	}
}

// Build builds a new format function.
//
// Returns:
//   - FormatFn: The new format function. Never returns nil.
func (b Builder) Build() FormatFn {
	var prefix rune

	if b.prefix == 0 {
		prefix = DefaultPrefix
	} else {
		prefix = b.prefix
	}

	idx := slices.Index(b.allowed_verbs, prefix)
	if idx != -1 {
		b.allowed_verbs = slices.Delete(b.allowed_verbs, idx, idx+1)
	}

	allowed_verbs := make([]rune, 0, len(b.allowed_verbs))

	for _, v := range b.allowed_verbs {
		allowed_verbs = append(allowed_verbs, v)
	}

	lexer := internal.NewLexer(prefix, allowed_verbs)

	fn := func(format string, data Formatter) (string, error) {
		if format == "" {
			return "", nil
		}

		defer lexer.Reset()

		var buff bytes.Buffer

		buff.WriteString(format)

		lexer.SetInputStream(&buff)

		err := lexer.Lex()
		if err != nil {
			return "", err
		}

		tokens := lexer.Tokens()

		res, err := apply(tokens, data)
		if err != nil {
			return "", err
		}

		return res, nil
	}

	return fn
}

// Reset resets the builder to make it reusable.
func (b *Builder) Reset() {
	if b == nil {
		return
	}

	b.prefix = 0

	if len(b.allowed_verbs) > 0 {
		b.allowed_verbs = b.allowed_verbs[:0]
	}
}
