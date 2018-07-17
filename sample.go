package main

import (
	"encoding"
	"math"
	"time"
)

// Sample represents any object for a complete sample of data
// This could be a discrete sensor reading, a set of readings
// (e.g. all the analog inputs), a collection of readings for
// a time slice or even a complete set of readings for a whole
// lap
type Sample interface {
	// Type returns the Sample Type
	Type() string
}

// ParseableSample is a sample that implements the BinaryUnmarshaler
// interface
type ParseableSample interface {
	Sample
	encoding.BinaryUnmarshaler
}

// LapMarker is a position to indicate the beginning of a sector
type LapMarker struct {
	// Marker is the marker number
	Marker int

	// Latitude is the north/south position of the marker in degrees
	Latitude Coordinate

	// Longitude is the east/west position of the marker in degrees
	Longitude Coordinate

	// Heading is the direction the marker is pointing
	Heading Heading
}

// Type returns the Sample Type, in this case "LapMarker"
func (*LapMarker) Type() string { return "LapMarker" }

// UnmarshalBinary parses the input byte buffer and assigns
// the parsed values to the LapMarker. BufError is returned
// if the input buffer is too short to process
func (lm *LapMarker) UnmarshalBinary(buf []byte) error {
	err := checkBufLen(buf, 16)
	if err == nil {
		lm.Marker = int(buf[0])
		lm.Latitude = Coordinate(computeGeo(buf[1:5])) * 0.0000001
		lm.Longitude = Coordinate(computeGeo(buf[6:10])) * 0.0000001
		lm.Heading = Heading(computeGeo(buf[11:15])) * 0.00001
	}
	println(hexDump(buf))
	return err
}

// LoggerStorage represents data received on the Logger Storage channel
// (data channel 6)
type LoggerStorage struct {
	// SerialNumber is the serial number of the data logger
	SerialNumber uint16

	// SoftwareVersion is the software version of the data logger
	SoftwareVersion int

	// BootloadVersion is the bootload version of the data logger
	BootloadVersion int
}

// Type returns the Sample Type, in this case "LoggerStorage"
func (*LoggerStorage) Type() string { return "LoggerStorage" }

// UnmarshalBinary parses the input byte buffer and assigns
// the parsed values to the LoggerStorage. BufError is returned
// if the input buffer is too short to process
func (ls *LoggerStorage) UnmarshalBinary(buf []byte) error {
	err := checkBufLen(buf, 4)
	if err == nil {
		ls.SerialNumber = uint16(buf[0])<<8 | uint16(buf[1])
		ls.SoftwareVersion = int(buf[2])
		ls.BootloadVersion = int(buf[3])
	}
	return err
}

// GPSTimeStorage contains the data received on the GPS Time Storage channel
// (data channel 7)
type GPSTimeStorage struct {
	// Time is the GPS Time - number if milliseconds since the beginning of the
	// week (midnight between Saturday and Sunday)
	Time GPSTime
}

// Type returns the Sample Type, in this case "GPSTimeStorage"
func (*GPSTimeStorage) Type() string { return "GPSTimeStorage" }

// UnmarshalBinary parses the input byte buffer and assigns
// the parsed values to the GPSTimeStorage. BufError is returned
// if the input buffer is too short to process
func (time *GPSTimeStorage) UnmarshalBinary(buf []byte) error {
	err := checkBufLen(buf, 4)
	if err == nil {
		time.Time = GPSTime(buf[0])<<24 | GPSTime(buf[1])<<16 | GPSTime(buf[2])<<8 | GPSTime(buf[3])
	}
	return err
}

// Accelerations contains the accelerometer data received on
// data channel 8
type Accelerations struct {
	// Lateral acceleration in G
	Lateral Acceleration

	// Longitudinal acceleration in G
	Longitudinal Acceleration
}

// Type returns the Sample Type, in this case "Accelerations"
func (*Accelerations) Type() string { return "Accelerations" }

// UnmarshalBinary parses the input byte buffer and assigns
// the parsed values to the Accelerations. BufError is returned
// if the input buffer is too short to process
func (accel *Accelerations) UnmarshalBinary(buf []byte) error {
	err := checkBufLen(buf, 4)
	if err == nil {
		accel.Lateral = computeAcceleration(buf[0:2])
		accel.Longitudinal = computeAcceleration(buf[2:4])
	}
	return err
}

// Vector is the square root of the sum of squares for lateral
// and longitudinal acceleration thus producing an acceleration
// vector in the X and Y axis
func (accel *Accelerations) Vector() Acceleration {
	lat := float64(accel.Lateral)
	long := float64(accel.Longitudinal)
	return Acceleration(math.Sqrt(lat*lat + long*long))
}

// Timestamp contains information received in data channel
// 9
type Timestamp struct {
	// Timestamp is the offset in milliseconds since the logging
	// session started
	Timestamp TimeOffset
}

// Type returns the Sample Type, in this case "TimeStamp"
func (*Timestamp) Type() string { return "Timestamp" }

// UnmarshalBinary parses the input byte buffer and assigns
// the parsed values to the Timestamp. BufError is returned
// if the input buffer is too short to process
func (ts *Timestamp) UnmarshalBinary(buf []byte) error {
	err := checkBufLen(buf, 3)
	if err == nil {
		value := int(buf[0])<<16 | int(buf[1])<<8 | int(buf[2])
		// convert centiseconds to milliseconds
		ts.Timestamp = TimeOffset(value) * 10
	}
	return err
}

// GPSPosition contains information sent on data channel 10
type GPSPosition struct {
	// Latitude (north/south) position in degrees
	Latitude Coordinate

	// Longitude (east/west) position in degrees
	Longitude Coordinate

	// Accuracy in millimeters
	Accuracy GPSAccuracy
}

// Type returns the Sample Type, in this case "GPSPosition"
func (*GPSPosition) Type() string { return "GPSPosition" }

// UnmarshalBinary parses the input byte buffer and assigns
// the parsed values to the GPSPosition. BufError is returned
// if the input buffer is too short to process
func (gpsPos *GPSPosition) UnmarshalBinary(buf []byte) error {
	err := checkBufLen(buf, 12)
	if err == nil {
		gpsPos.Latitude = Coordinate(computeGeo(buf[0:4])) * 0.0000001
		gpsPos.Longitude = Coordinate(computeGeo(buf[4:8])) * 0.0000001
		gpsPos.Accuracy = GPSAccuracy(buf[8])<<24 | GPSAccuracy(buf[9])<<16 | GPSAccuracy(buf[10])<<8 | GPSAccuracy(buf[11])
	}
	return err
}

// SpeedData contains the information received in data channel 11
type SpeedData struct {
	// Speed in meters/second
	Speed Speed

	// Accuracy in meters/second
	Accuracy SpeedAccuracy
}

// Type returns the Sample Type, in this case "SpeedData"
func (*SpeedData) Type() string { return "SpeedData" }

// UnmarshalBinary parses the input byte buffer and assigns
// the parsed values to the SpeedData. BufError is returned
// if the input buffer is too short to process
func (speed *SpeedData) UnmarshalBinary(buf []byte) error {
	err := checkBufLen(buf, 8)
	if err == nil {
		value := int(buf[0])<<24 | int(buf[1])<<16 | int(buf[2])<<8 | int(buf[3])
		speed.Speed = Speed(value) * 0.01
		value = int(buf[4])<<24 | int(buf[5])<<16 | int(buf[6])<<8 | int(buf[7])
		speed.Accuracy = SpeedAccuracy(value)
	}
	return err
}

// BeaconPulse is also sometimes referred to as Padding and is
// sent on data channel 12
type BeaconPulse struct {
	// Data is any data in the padding message
	Data byte
}

// Type returns the Sample Type, in this case "BeaconPulse"
func (*BeaconPulse) Type() string { return "BeaconPulse" }

// UnmarshalBinary parses the input byte buffer and assigns
// the parsed values to the BeaconPulse. BufError is returned
// if the input buffer is too short to process
func (beacon *BeaconPulse) UnmarshalBinary(buf []byte) error {
	err := checkBufLen(buf, 1)
	if err == nil {
		beacon.Data = buf[0]
	}
	return err
}

// FrequencyInput contains information sent on data channels
// 14-17 representing measured values on the frequency
// input pins of the data logger
type FrequencyInput struct {
	// Channel that was reported (FrequencyChannel1-FrequencyChannel5)
	Channel Channel

	// Frequency reported in hertz
	Frequency Frequency
}

// Type returns the Sample Type, in this case "FrequencyInput"
func (*FrequencyInput) Type() string { return "FrequencyInput" }

// UnmarshalBinary parses the input byte buffer and assigns
// the parsed values to the FrequencyInput. BufError is returned
// if the input buffer is too short to process
func (freq *FrequencyInput) UnmarshalBinary(buf []byte) error {
	err := checkBufLen(buf, 3)
	if err == nil {
		// number of 6 Mhz pulses received for the input signal
		value := int64(buf[0])<<16 | int64(buf[1])<<8 | int64(buf[2])

		// time (in seconds) taken for input signal
		time := float64(value) * 1.66666666666667E-07

		// frequency (in herz - 1s / period)
		freq.Frequency = Frequency(1 / time)
	}
	return err
}

// AnalogInput contains voltage readings from the analog inputs
// on the data logger (data channels 20-51)
type AnalogInput struct {
	// Channel the value is reported for (AnalogChannel1-AnalogChannel32)
	Channel Channel

	// Voltage reported from the sensor
	Voltage Voltage
}

// Type returns the Sample Type, in this case "AnalogInput"
func (*AnalogInput) Type() string { return "AnalogInput" }

// UnmarshalBinary parses the input byte buffer and assigns
// the parsed values to the AnalogInput. BufError is returned
// if the input buffer is too short to process
func (analog *AnalogInput) UnmarshalBinary(buf []byte) error {
	err := checkBufLen(buf, 2)
	if err == nil {
		value := int(buf[0])<<8 | int(buf[1])
		analog.Voltage = Voltage(value)
	}
	return err
}

// DateStorage contains the date and time received in data channel
// 55
type DateStorage struct {
	// Time is the current time sent by the data logger
	Time time.Time
}

// Type returns the Sample Type, in this case "DateStorage"
func (*DateStorage) Type() string { return "DateStorage" }

// UnmarshalBinary parses the input byte buffer and assigns
// the parsed values to the DateStorage. BufError is returned
// if the input buffer is too short to process
func (date *DateStorage) UnmarshalBinary(buf []byte) error {
	err := checkBufLen(buf, 8)
	if err == nil {
		offset := int(buf[7])
		// two's complement conversion
		if buf[7]&0x80 == 0x80 {
			offset = -1 * (int(^buf[7]) + 1)
		}

		// offset is in 15 minute increments, so convert this to seconds for
		// the timezone lookup
		offset = offset * 15 * 60
		zone := zoneLookup(offset)
		date.Time = time.Date(int(buf[5])<<8|int(buf[6]), time.Month(buf[4]), int(buf[3]), int(buf[2]), int(buf[1]), int(buf[0]), 0, zone)
	}
	return err
}

// CourseData contains the Heading information sent in data channel
// 56
type CourseData struct {
	// Heading is the direction the vehicle is moving, in degrees
	Heading Heading

	// Accuracy is the accuracy of the heading measurement in degrees
	Accuracy HeadingAccuracy
}

// Type returns the Sample Type, in this case "CourseData"
func (*CourseData) Type() string { return "CourseData" }

// UnmarshalBinary parses the input byte buffer and assigns
// the parsed values to the CourseData. BufError is returned
// if the input buffer is too short to process
func (course *CourseData) UnmarshalBinary(buf []byte) error {
	err := checkBufLen(buf, 8)
	if err == nil {
		course.Heading = Heading(computeGeo(buf[0:5])) * 0.00001
		course.Accuracy = HeadingAccuracy(computeGeo(buf[5:8])) * 0.00001
	}
	return err
}

// GPSAltitude is the altitude measurement reported by the GPS in the
// data logger (data channel 57)
type GPSAltitude struct {
	// Altitude is the distance above sea level measured in millimeters
	Altitude Altitude

	// Accuracy is the accuracy of the altitude measurement in millimeters
	Accuracy AltitudeAccuracy
}

// Type returns the Sample Type, in this case "GPSAltitude"
func (*GPSAltitude) Type() string { return "GPSAltitude" }

// UnmarshalBinary parses the input byte buffer and assigns
// the parsed values to the GPSAltitude. BufError is returned
// if the input buffer is too short to process
func (altitude *GPSAltitude) UnmarshalBinary(buf []byte) error {
	err := checkBufLen(buf, 8)
	if err == nil {
		altitude.Altitude = Altitude(buf[0])<<24 | Altitude(buf[1])<<16 | Altitude(buf[2])<<8 | Altitude(buf[3])
		altitude.Accuracy = AltitudeAccuracy(buf[4])<<24 | AltitudeAccuracy(buf[5])<<16 | AltitudeAccuracy(buf[6])<<8 | AltitudeAccuracy(buf[7])
	}
	return err
}

func sampleFactory(channel Channel, buf []byte) (sample ParseableSample) {
	switch {
	case channel == LapMarkerChannel:
		sample = &LapMarker{}
	case channel == LoggerStorageChannel:
		sample = &LoggerStorage{}
	case channel == GPSTimeStorageChannel:
		sample = &GPSTimeStorage{}
	case channel == AccelerationsChannel:
		sample = &Accelerations{}
	case channel == TimestampChannel:
		sample = &Timestamp{}
	case channel == GPSPositionChannel:
		sample = &GPSPosition{}
	case channel == SpeedDataChannel:
		sample = &SpeedData{}
	case channel == BeaconPulseChannel:
		sample = &BeaconPulse{}
	case FrequencyChannel1 <= channel && channel <= FrequencyChannel5:
		sample = &FrequencyInput{Channel: channel}
	case AnalogChannel1 <= channel && channel <= AnalogChannel32:
		sample = &AnalogInput{Channel: channel}
	case channel == DateStorageChannel:
		sample = &DateStorage{}
	case channel == CourseDataChannel:
		sample = &CourseData{}
	case channel == GPSAltitudeChannel:
		sample = &GPSAltitude{}
	case channel == RunStatusChannel:
		if len(buf) > 0 {
			if buf[0] <= 4 {
				sample = &StartStopInfo{}
			} else if buf[0] == 5 {
				sample = &TrackMarkerFailureMessage{}
			} else if buf[0] == 6 {
				sample = &Status{}
			}
		}
	}
	return
}
