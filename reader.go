package main

import (
	"bufio"
	"io"
)

// RunReader reads a run file and outputs the data channel
// messages
type RunReader struct {
	reader   *bufio.Reader
	messages chan Sample
}

// NewRunReader takes an io.Reader object and returns a
// RunReader ready to process messages from the underlying
// Reader
func NewRunReader(reader io.Reader) *RunReader {
	return &RunReader{
		reader:   bufio.NewReader(reader),
		messages: make(chan Sample),
	}
}

func (rr *RunReader) emit(message *Message) {
	rr.messages <- message
}

func (rr *RunReader) readChannel() error {
	b, err := rr.reader.ReadByte()
	if err == nil {
		channel := Channel(b)
		length := channelLengths[channel] - 1

		if length < 0 {
			return ErrUnknownLength
		}

		buf := make([]byte, length)
		_, err = io.ReadFull(rr.reader, buf)
		if err == nil {
			msg := &Message{Channel: channel, Data: buf[0 : length-1], Checksum: buf[length-1]}
			if msg.Valid() {
				rr.emit(msg)
			} else {
				msg = nil
				err = ErrChecksum
			}
		}
	}
	return err
}

func (rr *RunReader) synchronize() error {
	found := 0
	err := rr.readChannel()
	// look for two frames in a row that have a correct checksum
	for err == ErrUnknownLength || err == ErrChecksum || (err == nil && found < 2) {
		if err == nil {
			found++
		} else if err == ErrUnknownLength || err == ErrChecksum {
			found = 0
		}
		err = rr.readChannel()
	}
	return err
}

func (rr *RunReader) Read() error {
	err := rr.synchronize()

	for err == nil {
		err = rr.readChannel()
	}

	if err == ErrChecksum {
		var b byte
		b, err = rr.reader.ReadByte()
		if err == nil {
			length := channelLengths[Channel(b)]
			io.ReadFull(rr.reader, make([]byte, length))
			for b, err = rr.reader.ReadByte(); err == nil; b, err = rr.reader.ReadByte() {
				if b != 0x00 {
					return ErrUnexpectedInput
				}
			}
		}
	}
	close(rr.messages)
	return err
}

// Output returns the channel that Samples can be read from
func (rr *RunReader) Output() <-chan Sample {
	return rr.messages
}
