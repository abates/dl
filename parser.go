package main

import (
	"fmt"
)

// RunParser is used to read data messages from the data loger
// and parse them into data samples to be used by a demuxer or
// other downstream processor
type RunParser struct {
}

// Process accepts a sample message and unmarshals the data specific data types
func (rp *RunParser) Process(input <-chan Sample, output chan<- Sample) {
	for sample := range input {
		if message, ok := sample.(*Message); ok {
			newSample := sampleFactory(message.Channel, message.Data)
			if newSample == nil {
				if message.Channel != RunInformationChannel {
					fmt.Printf("Message: %v\n", message)
				}
				output <- sample
			} else {
				newSample.UnmarshalBinary(message.Data)
				output <- newSample
			}
		}
	}
	close(output)
}
