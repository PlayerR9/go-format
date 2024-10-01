package f_string

import (
	"io"

	"github.com/PlayerR9/go-errors/assert"
	"github.com/PlayerR9/go-format/f_string/internal"
)

// FStringer is the interface for a format string traverser.
type FStringer interface {
	// FString traverses the format string.
	//
	// Parameters:
	//   - trav: The traverser to use. Assumed to be non-nil.
	//
	// Returns:
	//   - error: An error if the format string could not be traversed.
	FString(trav Traverser) error
}

// New creates a new traverser.
//
// Parameters:
//   - w: The writer to write to.
//
// Returns:
//   - *Traversor: The new traverser. Never returns nil.
//
// If w is nil, it is set to io.Discard.
func New(w io.Writer) *Traversor {
	if w == nil {
		w = io.Discard
	}

	buffer, err := internal.NewBuffer(w)
	assert.Err(err, "internal.NewBuffer(w)")

	return &Traversor{
		buffer: buffer,
	}
}
