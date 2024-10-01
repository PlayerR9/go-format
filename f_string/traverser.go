package f_string

import (
	"fmt"
	"io"
	"strings"

	gers "github.com/PlayerR9/go-errors"
	"github.com/PlayerR9/go-errors/assert"
	"github.com/PlayerR9/go-format/f_string/internal"
)

// Traverser is the interface for a format string traverser.
type Traverser interface {
	// WriteLine writes a line to the traverser.
	//
	// Parameters:
	//   - str: The line to write.
	//
	// Returns:
	//   - error: An error if the line could not be written.
	WriteLine(str string) error

	io.Writer
}

// Traversor is the traverser.
type Traversor struct {
	// buffer is the buffer.
	buffer *internal.Buffer

	// indent_level is the indent level.
	indent_level int
}

// WriteLine writes a line to the traverser.
//
// Parameters:
//   - str: The line to write.
//
// Returns:
//   - error: An error if the line could not be written.
func (t *Traversor) WriteLine(str string) error {
	if t == nil {
		err := gers.NewErrNilReceiver()
		err.AddFrame("*Traversor.WriteLine()")

		return err
	}

	if str == "" {
		err := t.buffer.WriteEmptyLine()
		if err != nil {
			err.AddFrame("*Traversor.WriteLine()")

			return err
		}
	} else {
		indent := strings.Repeat("   ", t.indent_level)

		line := fmt.Sprintf("%s%s", indent, str)

		err := t.buffer.WriteLine(line)
		if err != nil {
			err.AddFrame("*Traversor.WriteLine()")

			return err
		}
	}

	return nil
}

// Write writes data to the traverser.
//
// Parameters:
//   - data: The data to write.
//
// Returns:
//   - int: The number of bytes written.
//   - error: An error if the data could not be written.
//
// Always returns len(data) when error is nil.
func (t *Traversor) Write(data []byte) (int, error) {
	if len(data) == 0 {
		return 0, nil
	}

	if t == nil {
		err := gers.NewErrNilReceiver()
		err.AddFrame("*Traversor.Write()")

		return 0, err
	}

	assert.NotNil(t.buffer, "t.buffer")

	err := t.buffer.Write(data)
	if err != nil {
		err.AddFrame("*Traversor.Write()")

		return 0, err
	}

	return len(data), nil
}

/* func (t *Traversor) AcceptLine() {
	if t == nil || t.builder.Len() == 0 {
		return
	}

	t.lines = append(t.lines, t.builder.String())
	t.builder.Reset()
}
*/

// IndentBy increments the indent level by n.
//
// Parameters:
//   - n: The amount to increment the indent level by.
//
// Returns:
//   - *Traversor: The traverser.
//   - error: An error if the traverser could not be created.
//
// If the new indent level is less than 0, it is set to 0.
func (t *Traversor) IndentBy(n int) (*Traversor, error) {
	if t == nil {
		err := gers.NewErrNilReceiver()
		err.AddFrame("*Traversor.IndentBy()")

		return nil, err
	}

	new_indent_level := t.indent_level + n
	if new_indent_level < 0 {
		new_indent_level = 0
	}

	return &Traversor{
		buffer:       t.buffer,
		indent_level: new_indent_level,
	}, nil
}
