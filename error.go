package main

import (
	"fmt"
	"io"
)

var (
	// ErrShortBuffer is indicated as an error cause when not enough bytes
	// are available in an input buffer
	ErrShortBuffer = fmt.Errorf("Input buffer was short")

	// ErrChecksum is indicated when a message has an invalid checksum
	ErrChecksum = fmt.Errorf("Checksum Error")

	// ErrUnknownLength indicates that the message length is not known for the
	// input data channel
	ErrUnknownLength = fmt.Errorf("Unknown message length")

	// ErrEOF indicates no more data is available for reading
	ErrEOF = io.EOF

	// ErrUnexpectedInput indicates an unknown or otherwise unexpected value
	// was received in a byte stream
	ErrUnexpectedInput = fmt.Errorf("Unexpected Input")
)

// BufError has information to indicate a buffer error
type BufError struct {
	Cause error // the underlying error
	Need  int   // the number of bytes required
	Got   int   // the length of the supplied buffer
}

func newBufError(cause error, need int, got int) *BufError {
	return &BufError{Cause: cause, Need: need, Got: got}
}

// Error will indicate what caused a buffer error to occur
func (be *BufError) Error() string {
	cause := ""
	if be.Cause != nil {
		cause = fmt.Sprintf("%v: ", be.Cause)
	}
	return fmt.Sprintf("%sneed %d bytes got %d", cause, be.Need, be.Got)
}

// ParseError indicates a failure in parsing a byte buffer
type ParseError struct {
	reason string
}

func newParseError(reason string) error {
	return &ParseError{reason}
}

// Error returns the reason for the parse error
func (pe *ParseError) Error() string {
	return pe.reason
}

func checkBufLen(buf []byte, need int) error {
	if len(buf) < need {
		return newBufError(ErrShortBuffer, need, len(buf))
	}
	return nil
}
