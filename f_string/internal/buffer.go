package internal

import (
	"io"

	gers "github.com/PlayerR9/go-errors"
	"github.com/PlayerR9/go-errors/assert"
)

// Buffer is a buffer for writing to an io.Writer.
type Buffer struct {
	// w is the writer to write to.
	w io.Writer
}

// NewBuffer creates a new buffer.
//
// Parameters:
//   - w: The writer to write to.
//
// Returns:
//   - *Buffer: The new buffer.
//   - *error: An error if the buffer could not be created.
func NewBuffer(w io.Writer) (*Buffer, error) {
	if w == nil {
		err := gers.NewErrNilReceiver("internal.NewBuffer()")

		return nil, err
	}

	return &Buffer{
		w: w,
	}, nil
}

// Write writes data to the buffer.
//
// Parameters:
//   - data: The data to write.
//
// Returns:
//   - *errors.Err: An error if the data could not be written.
func (b *Buffer) Write(data []byte) *gers.Err {
	assert.NotNil(b, "receiver")

	n, err := b.w.Write(data)
	if err != nil {
		err := gers.NewFromError(gers.OperationFail, err)
		err.AddFrame("*Buffer.Write()")

		return err
	} else if n != len(data) {
		err := gers.NewFromError(gers.OperationFail, io.ErrShortWrite)
		err.AddFrame("*Buffer.Write()")

		return err
	}

	return nil
}

// WriteLine writes a line to the buffer.
//
// Parameters:
//   - str: The line to write.
//
// Returns:
//   - *errors.Err: An error if the line could not be written.
func (b *Buffer) WriteLine(str string) *gers.Err {
	assert.NotNil(b, "receiver")

	str = str + "\n"

	data := []byte(str)

	n, err := b.w.Write(data)
	if err != nil {
		err := gers.NewFromError(gers.OperationFail, err)
		err.AddFrame("*Buffer.WriteLine()")

		return err
	}

	if n != len(data) {
		err := gers.NewFromError(gers.OperationFail, io.ErrShortWrite)
		err.AddFrame("*Buffer.WriteLine()")

		return err
	}

	return nil
}

// WriteEmptyLine writes an empty line to the buffer.
//
// Returns:
//   - *errors.Err: An error if the line could not be written.
func (b *Buffer) WriteEmptyLine() *gers.Err {
	assert.NotNil(b, "receiver")

	data := []byte("\n")

	n, err := b.w.Write(data)
	if err != nil {
		err := gers.NewFromError(gers.OperationFail, err)
		err.AddFrame("*Buffer.WriteEmptyLine()")

		return err
	}

	if n != len(data) {
		err := gers.NewFromError(gers.OperationFail, io.ErrShortWrite)
		err.AddFrame("*Buffer.WriteEmptyLine()")

		return err
	}

	return nil
}
