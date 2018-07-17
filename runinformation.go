//go:generate stringer -type=StartMethod,StopMethod,PreTriggerLoopMethod,AutoMethod,TrackMarkerFailureCode
package main

import (
	"fmt"

	"gonum.org/v1/gonum/unit"
)

// StartMethod indicates how the logging session was started
type StartMethod int

// Logging session start methods
const (
	ButtonStart StartMethod = iota + 1
	AutoStart
	PreTriggerLoop
)

// StopMethod indicates how the logging session was started
type StopMethod int

// Logging session stop methods
const (
	Running StopMethod = iota
	ButtonStop
	AutoStop
	PostTriggerExpired
	LowBatteryVoltage
	LowBufferSpace
	GSMCommand
	DiskFull
	DVRSerialCommand
	RemoteRTSerialCommand
	FileLimitExceeded
	LoopSizeExceeded
	CAMDisconnect
)

// PreTriggerLoopMethod indicates how the pre-trigger loop was started
type PreTriggerLoopMethod int

// Pre-trigger loop methods
const (
	ButtonPress PreTriggerLoopMethod = iota + 1
	Autostart
)

// AutoMethod indicates how the logging session was auto-started
type AutoMethod int

// Automatic starting methods
const (
	ADC1 AutoMethod = iota + 1
	ADC2
	ADC3
	ADC4
	ADC5
	ADC6
	ADC7
	ADC8

	LateralGForce      AutoMethod = 15
	LongitudinalGForce AutoMethod = 16
)

// StartStopInfo is a status message received on data channel 2
type StartStopInfo struct {
	// StartMethod for the logging session
	StartMethod StartMethod

	// StopMethod for the logging session
	StopMethod StopMethod

	// PreTriggerLoopMethod for the logging session
	PreTriggerLoopMethod PreTriggerLoopMethod

	// PreTriggerTie for the logging session
	PreTriggerTime unit.Time

	// PostTriggerTime for the logging session
	PostTriggerTime unit.Time

	// AutoStartSource for the logging session
	AutoStartSource AutoMethod

	// AutoStopSource for the logging session
	AutoStopSource AutoMethod

	// LowestBuffer for the logging session
	LowestBuffer int
}

// Type returns the Sample Type, in this case "StartStopInfo"
func (*StartStopInfo) Type() string { return "StartStopInfo" }

// UnmarshalBinary parses the input byte buffer and assigns
// the parsed values to the StartStopInfo. BufError is returned
// if the input buffer is too short to process.  Additionally,
// a ParseError is returned if the start method cannot be determined
func (ssm *StartStopInfo) UnmarshalBinary(buf []byte) (err error) {
	err = checkBufLen(buf, 9)
	if err == nil {
		if buf[0] <= 0x03 {
			ssm.StartMethod = StartMethod(buf[0])
			ssm.StopMethod = StopMethod(buf[1])
		} else if buf[0] == 0x04 {
			ssm.StartMethod = StartMethod(buf[1] >> 4)
			ssm.StopMethod = StopMethod(buf[1] & 0x0f)
		} else {
			err = newParseError(fmt.Sprintf("Unknown start method 0x%02x", buf[0]))
		}
	}

	if err == nil {
		ssm.PreTriggerLoopMethod = PreTriggerLoopMethod(buf[2])
		ssm.PreTriggerTime = unit.Time(buf[3]) * unit.Centisecond
		ssm.PostTriggerTime = unit.Time(buf[4]) * unit.Centisecond
		ssm.AutoStartSource = AutoMethod(buf[5])
		ssm.AutoStopSource = AutoMethod(buf[6])
		ssm.LowestBuffer = int(buf[7])<<8 | int(buf[8])
	}
	return err
}

// TrackMarkerFailureCode indicates the reason the data logger failed to add
// a track marker
type TrackMarkerFailureCode int

// Possible track marker failure reasons
const (
	MarkerAlreadExists TrackMarkerFailureCode = iota + 1
	TooManyMarkers
	NotLoggingData
	NoCard
	GPSInnacurate
	CorrectGPSNotAvailable
)

// TrackMarkerFailureMessage is a status message indicating that the data
// logger failed to add a track marker, as requested. This message is
// received on data channel 2 (Run Status Messages)
type TrackMarkerFailureMessage struct {
	// Code indicates why the data logger failed to add a track marker
	Code TrackMarkerFailureCode
}

// Type returns the Sample Type, in this case "TrackMarkerFailureMessage"
func (*TrackMarkerFailureMessage) Type() string { return "TrackMarkerFailureMessage" }

// UnmarshalBinary parses the input byte buffer and assigns
// the parsed values to the TrackMarkerFailureMessage. BufError is returned
// if the input buffer is too short to process
func (tmfm *TrackMarkerFailureMessage) UnmarshalBinary(buf []byte) (err error) {
	err = checkBufLen(buf, 2)
	if err == nil {
		tmfm.Code = TrackMarkerFailureCode(buf[1])
	}
	return err
}

// Status contains information about the data logger status. Status
// messages are received on data channel 2 (Run Status Messages)
type Status struct {
	// GPSDetected indicates if the GPS module was detected
	GPSDetected bool

	// IMUDetected indicates if the IMU is detected
	IMUDetected bool

	// GPS1Lock
	GPS1Lock bool

	// GPS2Lock
	GPS2Lock bool

	// CarrierLock
	CarrierLock bool

	// RTKLock
	RTKLock bool

	// INSInitialized
	INSInitialized bool

	// INSConverged
	INSConverged bool
}

// Type returns the Sample Type, in this case "Status"
func (*Status) Type() string { return "Status" }

// UnmarshalBinary parses the input byte buffer and assigns
// the parsed values to the Status. BufError is returned
// if the input buffer is too short to process
func (sm *Status) UnmarshalBinary(buf []byte) (err error) {
	sm.GPSDetected = buf[1]&0x80 == 0x80
	sm.IMUDetected = buf[1]&0x40 == 0x40
	sm.GPS1Lock = buf[1]&0x20 == 0x20
	sm.GPS2Lock = buf[1]&0x10 == 0x10
	sm.CarrierLock = buf[1]&0x08 == 0x08
	sm.RTKLock = buf[1]&0x04 == 0x04
	sm.INSInitialized = buf[1]&0x02 == 0x02
	sm.INSConverged = buf[1]&0x01 == 0x01
	return nil
}
