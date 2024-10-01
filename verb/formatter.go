package verb

import (
	"fmt"
	"strings"

	"github.com/PlayerR9/go-errors/assert"
	"github.com/PlayerR9/go-format/verb/internal"
)

// FormatFn is the type for a format function.
//
// Parameters:
//   - format: The format string.
//   - data: The data to format.
//
// Returns:
//   - string: The formatted string.
//   - error: An error if the format could not be formatted.
type FormatFn func(format string, data Formatter) (string, error)

// Formatter is the interface for a formatter.
type Formatter interface {
	// Format formats the data using the format string.
	//
	// Parameters:
	//   - format: The format string.
	//
	// Returns:
	//   - string: The formatted string.
	//   - error: An error if the format could not be formatted.
	Format(verb string) (string, error)
}

// apply is private function that applies the formatting to a series of tokens.
//
// Parameters:
//   - tokens: The list of tokens.
//   - data: The data to format.
//
// Returns:
//   - string: The formatted string.
//   - error: An error if the format could not be formatted.
func apply(tokens []*internal.Token, data Formatter) (string, error) {
	if len(tokens) == 0 {
		return "", nil
	}

	var builder strings.Builder

	if data == nil {
		for _, token := range tokens {
			assert.NotNil(token, "token")

			if token.IsVerb {
				return builder.String(), fmt.Errorf("verb %q is not supported", token.Data)
			}

			builder.WriteString(token.Data)
		}
	} else {
		for _, token := range tokens {
			assert.NotNil(token, "token")

			if !token.IsVerb {
				builder.WriteString(token.Data)
			} else {
				str, err := data.Format(token.Data)
				if err != nil {
					return builder.String(), err
				}

				builder.WriteString(str)
			}
		}
	}

	return builder.String(), nil
}
