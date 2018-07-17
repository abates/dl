package main

type Sector struct {
	epochs []Epoch
}

type SectorInfo struct {
	markers []LapMarker
}

func (si *SectorInfo) AddMarker(marker *LapMarker) {
	si.markers = append(si.markers, *marker)
}

type SectorAnalyzer struct {
	sectorInfo *SectorInfo
}

func NewSectorAnalyzer() *SectorAnalyzer { return &SectorAnalyzer{} }

func (sa *SectorAnalyzer) SectorInfo(sectorInfo *SectorInfo) *SectorAnalyzer {
	sa.sectorInfo = sectorInfo
	return sa
}

// Demux will start the demux loop.  This should usually be run in
// a go routine
func (sa *SectorAnalyzer) Process(input <-chan Sample, output chan<- Sample) {
	for range input {
	}
	close(output)
}
