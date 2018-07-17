package main

import (
	"fmt"
	"os"
	"runtime/pprof"
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

func NewPrintAnalyzer() *PrintAnalyzer { return &PrintAnalyzer{} }

func (pa *PrintAnalyzer) Process(input <-chan Sample, output chan<- Sample) {
	for sample := range input {
		fmt.Printf("%v\n", sample)
	}
	close(output)
}

func main() {
	f, err := os.Create("profile.pprof")
	if err != nil {
		panic("could not create CPU profile")
	}

	if err := pprof.StartCPUProfile(f); err != nil {
		panic("could not start CPU profile: ")
	}

	defer pprof.StopCPUProfile()

	input, err := os.Open("test.run")
	if err == nil {
		reader := NewRunReader(input)
		go reader.Read()
		processor := NewProcessingChain(reader.Output()).Append(&RunParser{}).Append(&SampleDemuxer{}).Append(NewPrintAnalyzer())
		processor.Wait()
	}

	if err != nil {
		fmt.Printf("Completed: %v\n", err)
	}
}
