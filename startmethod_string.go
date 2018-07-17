// Code generated by "stringer -type=StartMethod,StopMethod,PreTriggerLoopMethod,AutoMethod,TrackMarkerFailureCode"; DO NOT EDIT.

package main

import "strconv"

const _StartMethod_name = "ButtonStartAutoStartPreTriggerLoop"

var _StartMethod_index = [...]uint8{0, 11, 20, 34}

func (i StartMethod) String() string {
	i -= 1
	if i < 0 || i >= StartMethod(len(_StartMethod_index)-1) {
		return "StartMethod(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _StartMethod_name[_StartMethod_index[i]:_StartMethod_index[i+1]]
}

const _StopMethod_name = "RunningButtonStopAutoStopPostTriggerExpiredLowBatteryVoltageLowBufferSpaceGSMCommandDiskFullDVRSerialCommandRemoteRTSerialCommandFileLimitExceededLoopSizeExceededCAMDisconnect"

var _StopMethod_index = [...]uint8{0, 7, 17, 25, 43, 60, 74, 84, 92, 108, 129, 146, 162, 175}

func (i StopMethod) String() string {
	if i < 0 || i >= StopMethod(len(_StopMethod_index)-1) {
		return "StopMethod(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _StopMethod_name[_StopMethod_index[i]:_StopMethod_index[i+1]]
}

const _PreTriggerLoopMethod_name = "ButtonPressAutostart"

var _PreTriggerLoopMethod_index = [...]uint8{0, 11, 20}

func (i PreTriggerLoopMethod) String() string {
	i -= 1
	if i < 0 || i >= PreTriggerLoopMethod(len(_PreTriggerLoopMethod_index)-1) {
		return "PreTriggerLoopMethod(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _PreTriggerLoopMethod_name[_PreTriggerLoopMethod_index[i]:_PreTriggerLoopMethod_index[i+1]]
}

const (
	_AutoMethod_name_0 = "ADC1ADC2ADC3ADC4ADC5ADC6ADC7ADC8"
	_AutoMethod_name_1 = "LateralGForceLongitudinalGForce"
)

var (
	_AutoMethod_index_0 = [...]uint8{0, 4, 8, 12, 16, 20, 24, 28, 32}
	_AutoMethod_index_1 = [...]uint8{0, 13, 31}
)

func (i AutoMethod) String() string {
	switch {
	case 1 <= i && i <= 8:
		i -= 1
		return _AutoMethod_name_0[_AutoMethod_index_0[i]:_AutoMethod_index_0[i+1]]
	case 15 <= i && i <= 16:
		i -= 15
		return _AutoMethod_name_1[_AutoMethod_index_1[i]:_AutoMethod_index_1[i+1]]
	default:
		return "AutoMethod(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}

const _TrackMarkerFailureCode_name = "MarkerAlreadExistsTooManyMarkersNotLoggingDataNoCardGPSInnacurateCorrectGPSNotAvailable"

var _TrackMarkerFailureCode_index = [...]uint8{0, 18, 32, 46, 52, 65, 87}

func (i TrackMarkerFailureCode) String() string {
	i -= 1
	if i < 0 || i >= TrackMarkerFailureCode(len(_TrackMarkerFailureCode_index)-1) {
		return "TrackMarkerFailureCode(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _TrackMarkerFailureCode_name[_TrackMarkerFailureCode_index[i]:_TrackMarkerFailureCode_index[i+1]]
}
