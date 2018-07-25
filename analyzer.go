package dl

import (
	"fmt"
	"sync"
)

const BufLen = 512

type Analyzer interface {
	Process(<-chan Sample, chan<- Sample)
}

type ProcessingChain struct {
	sync.WaitGroup
	channels []<-chan Sample
}

func NewProcessingChain(input <-chan Sample) *ProcessingChain {
	return &ProcessingChain{
		channels: []<-chan Sample{input},
	}
}

func (pc *ProcessingChain) Append(analyzer Analyzer) *ProcessingChain {
	input := pc.channels[len(pc.channels)-1]
	output := make(chan Sample, BufLen)
	pc.channels = append(pc.channels, output)
	pc.WaitGroup.Add(1)
	go func() {
		analyzer.Process(input, output)
		pc.WaitGroup.Done()
	}()
	return pc
}

type PrintAnalyzer struct{}

func (pa *PrintAnalyzer) Process(input <-chan Sample, output chan<- Sample) {
	for sample := range input {
		fmt.Printf("%v\n", sample)
	}
	close(output)
}
