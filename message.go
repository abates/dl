package dl

import (
	"fmt"
	"strings"
)

// Message is one frame of data received from the data loger. Each
// message includes the channel the data was received for, the data
// received and the checksum byte sent by the data logger
type Message struct {
	// Channel is the data channel the data was received for
	Channel Channel

	// Data is the byte string sent for this particular data message
	Data []byte

	// Checksum is sum of all the bytes (minus the checksum byte) in the
	// message
	Checksum byte
}

// Type returns the Sample Type, in this case "Message"
func (message *Message) Type() string { return "Message" }

// Length returns the full message length:
// 1 byte for Channel
// 1 byte for Checksum
// n bytes for len(message.Data)
func (message *Message) Length() int {
	return len(message.Data) + 2
}

// Valid computes the checksum of the message's data and compares it
// to the expected checksum.  If the message's data is valid then
// Valid() returns true
func (message *Message) Valid() bool {
	checksum := byte(message.Channel)
	for i := 0; i < len(message.Data); i++ {
		checksum += message.Data[i]
	}

	return checksum == message.Checksum
}

func hexDump(buf []byte) string {
	str := make([]string, len(buf))
	for i, b := range buf {
		str[i] = fmt.Sprintf("0x%02x", b)
	}
	return strings.Join(str, ", ")
}

// String returns a string representation of the message
// TODO replace this with Format
func (message *Message) String() string {
	valid := "invalid"
	if message.Valid() {
		valid = "valid"
	}
	return fmt.Sprintf("%v %s (%s)", message.Channel, hexDump(message.Data), valid)
}
