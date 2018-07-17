package main

import (
	"time"
)

// Epoch contains all the sampled data for one time slice
type Epoch struct {
	// Time is the time received by the GPS
	Time time.Time

	// Start is the starting offset for this Epoch
	Start TimeOffset

	// Stop is the ending offset for this Epoch
	Stop TimeOffset

	// TimeSlip indicates whether the vehicle is going faster
	// or slower at ta given point.  TimeSlip is a cumulative
	// value and is reset at the Start/Finish lap marker.
	// TimeSlip is calculated by the LapAnalyzer and is not a measured
	// value
	TimeSlip TimeOffset

	// StartStopInfo is the current information regarding how the session started or stopped
	StartStopInfo StartStopInfo

	// GPSTime is the time in milliseconds since the beginning of the week
	GPSTime GPSTime

	// Speed is the current speed in m/s
	Speed Speed

	// LateralAcceleration is the current lateral acceleration in G. Positive
	// values indicate cornering around a right-hand turn, negative values
	// indicate cornering around a left-hand turn
	LateralAcceleration Acceleration

	// LongitudinalAcceleration is the current longitudinal acceleration in G.
	// Positive values indicate acceleration in speed and negative values
	// indicate braking
	LongitudinalAcceleration Acceleration

	// VectorAcceleration is the square root of the sum of squares of lateral
	// and longitudinal acceleration.
	VectorAcceleration Acceleration

	// Heading is the direction (in degrees) that the vehicle is moving
	Heading Heading

	// Latitude is the current latitude (in degrees) of the vehicle
	Latitude Coordinate

	// Longitude is the current longitude (in degrees) of the vehicle
	Longitude Coordinate

	// Altitude is the current altitude (in millimeters) of the vehicle
	Altitude Altitude

	// AnalogInputs is a map containing the voltages of reported analog inputs
	AnalogInputs map[Channel]Voltage

	// FrequencyInputs is a map containing the frequency (in hertz) of reported frequency inputs
	FrequencyInputs map[Channel]Frequency

	// HeadingAccuracy indidcates the accuracy of the heading in degrees
	HeadingAccuracy HeadingAccuracy

	// AltitudeAccuracy indidcates the accuracy of the altitude in millimeters
	AltitudeAccuracy AltitudeAccuracy

	// GPSAccuracy indicates the accuracy of the latitude and longitude in millimeters
	GPSAccuracy GPSAccuracy
}

// Type returns the Sample Type, in this case "Epoch"
func (epoch *Epoch) Type() string { return "Epoch" }
