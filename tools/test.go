package main

import (
	"fmt"
	"os"
	"runtime/pprof"

	"github.com/abates/dl"
)

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
		reader := dl.NewRunReader(input)
		go reader.Read()
		processor := dl.NewProcessingChain(reader.Output()).Append(&dl.RunParser{}).Append(&dl.SampleDemuxer{}).Append(&dl.PrintAnalyzer{})
		processor.Wait()
	}

	if err != nil {
		fmt.Printf("Completed: %v\n", err)
	}
}
