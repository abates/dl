// Copyright 2018 Andrew Bates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dl

import "fmt"

// TimeOffset is the elapsed time. TimeOffset is measured in milliseconds
type TimeOffset int64

// Format satisfies interface fmt.Formatter
func (time TimeOffset) Format(f fmt.State, c rune) { formatUnit(f, c, "ms", time, int64(time)) }

// Speed is the vehicle speed. Speed is measured in meters per second
type Speed float64

// Format satisfies interface fmt.Formatter
func (speed Speed) Format(f fmt.State, c rune) { formatUnit(f, c, "m/s", speed, float64(speed)) }

// SpeedAccuracy is the accuracy of the speed measurement. SpeedAccuracy is measured in millimeters per second
type SpeedAccuracy int

// Format satisfies interface fmt.Formatter
func (accuracy SpeedAccuracy) Format(f fmt.State, c rune) {
	formatUnit(f, c, "mm/s", accuracy, int(accuracy))
}

// Coordinate is a single point on either the lattitudinal or longitudinal axis. Coordinate is measured in degrees
type Coordinate float64

// Format satisfies interface fmt.Formatter
func (coordinate Coordinate) Format(f fmt.State, c rune) {
	formatUnit(f, c, "°", coordinate, float64(coordinate))
}

// GPSAccuracy is the accuracy of the GPS coordinates. GPSAccuracy is measured in millimeters
type GPSAccuracy int

// Format satisfies interface fmt.Formatter
func (accuracy GPSAccuracy) Format(f fmt.State, c rune) {
	formatUnit(f, c, "mm", accuracy, int(accuracy))
}

// Heading is the direction something is headed. In the case of course information the heading is the direction the vehicle is moving. For lap markers, the heading is direction the marker is pointing. Heading is measured in degrees
type Heading float64

// Format satisfies interface fmt.Formatter
func (heading Heading) Format(f fmt.State, c rune) { formatUnit(f, c, "°", heading, float64(heading)) }

// HeadingAccuracy is the accuracy of the GPS heading. HeadingAccuracy is measured in degrees
type HeadingAccuracy float64

// Format satisfies interface fmt.Formatter
func (accuracy HeadingAccuracy) Format(f fmt.State, c rune) {
	formatUnit(f, c, "°", accuracy, float64(accuracy))
}

// Acceleration is the acceleration of the vehicle in a given direction. Acceleration is measured in standard gravity
type Acceleration float64

// Format satisfies interface fmt.Formatter
func (acceleration Acceleration) Format(f fmt.State, c rune) {
	formatUnit(f, c, "G", acceleration, float64(acceleration))
}

// GPSTime is the number of milliseconds since midnight between Saturday and Sunday. GPSTime is measured in milliseconds
type GPSTime uint32

// Format satisfies interface fmt.Formatter
func (gpsTime GPSTime) Format(f fmt.State, c rune) { formatUnit(f, c, "ms", gpsTime, uint32(gpsTime)) }

// Voltage is a measurement sampled from the anlog inputs of the data logger. Voltage is measured in millivolts
type Voltage int

// Format satisfies interface fmt.Formatter
func (voltage Voltage) Format(f fmt.State, c rune) { formatUnit(f, c, "mV", voltage, int(voltage)) }

// Frequency is a measurement sampled from the frequency inputs of the data logger. Frequency is measured in hertz
type Frequency float64

// Format satisfies interface fmt.Formatter
func (freq Frequency) Format(f fmt.State, c rune) { formatUnit(f, c, "hz", freq, float64(freq)) }

// Altitude is the height above sea level as measured by GPS. Altitude is measured in millimeters
type Altitude int

// Format satisfies interface fmt.Formatter
func (altitude Altitude) Format(f fmt.State, c rune) { formatUnit(f, c, "mm", altitude, int(altitude)) }

// AltitudeAccuracy is the accuracy of the altitude measurement. AltitudeAccuracy is measured in millimeters
type AltitudeAccuracy int

// Format satisfies interface fmt.Formatter
func (accuracy AltitudeAccuracy) Format(f fmt.State, c rune) {
	formatUnit(f, c, "mm", accuracy, int(accuracy))
}
