package main

// SampleDemuxer will read a parse stream and demultiplex the samples into
// time slices. These timeslices will be represented as Epochs
type SampleDemuxer struct {
}

func copyEpoch(input *Epoch) *Epoch {
	output := &Epoch{}
	*output = *input
	output.AnalogInputs = make(map[Channel]Voltage)
	for k, v := range input.AnalogInputs {
		output.AnalogInputs[k] = v
	}

	output.FrequencyInputs = make(map[Channel]Frequency)
	for k, v := range input.FrequencyInputs {
		output.FrequencyInputs[k] = v
	}

	return output
}

// Demux will start the demux loop.  This should usually be run in
// a go routine
func (dm *SampleDemuxer) Process(input <-chan Sample, output chan<- Sample) {
	epoch := &Epoch{
		AnalogInputs:    make(map[Channel]Voltage),
		FrequencyInputs: make(map[Channel]Frequency),
	}
	for sample := range input {
		switch v := sample.(type) {
		case *GPSTimeStorage:
			epoch.GPSTime = v.Time
		case *Accelerations:
			epoch.LateralAcceleration = v.Lateral
			epoch.LongitudinalAcceleration = v.Longitudinal
			epoch.VectorAcceleration = v.Vector()
		case *Timestamp:
			epoch.Stop = v.Timestamp
			output <- copyEpoch(epoch)
			epoch.Start = v.Timestamp
		case *GPSPosition:
			epoch.Latitude = v.Latitude
			epoch.Longitude = v.Longitude
			epoch.GPSAccuracy = v.Accuracy
		case *SpeedData:
			epoch.Speed = v.Speed
		case *FrequencyInput:
			epoch.FrequencyInputs[v.Channel] = v.Frequency
		case *AnalogInput:
			epoch.AnalogInputs[v.Channel] = v.Voltage
		case *DateStorage:
			epoch.Time = v.Time
		case *CourseData:
			epoch.Heading = v.Heading
			epoch.HeadingAccuracy = v.Accuracy
		case *GPSAltitude:
			epoch.Altitude = v.Altitude
			epoch.AltitudeAccuracy = v.Accuracy
		case *StartStopInfo:
			epoch.StartStopInfo = *v
		default:
			output <- sample
		}
	}
	close(output)
}
