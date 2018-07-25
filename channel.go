//go:generate stringer -type=Channel -linecomment=true
package dl

// Channel is the data channel that a data frame (in the serial stream) represents
type Channel int

// All the known data channels
const (
	RunInformationChannel             Channel = 1   // Run Information
	RunStatusChannel                  Channel = 2   // Run Status
	NewSectorTime                     Channel = 4   // New Sector Time
	LapMarkerChannel                  Channel = 5   // Lap Marker
	LoggerStorageChannel              Channel = 6   // Logger Storage
	GPSTimeStorageChannel             Channel = 7   // GPS Time Storage
	AccelerationsChannel              Channel = 8   // Accelerations
	TimestampChannel                  Channel = 9   // Timestamp
	GPSPositionChannel                Channel = 10  // GPS Position
	SpeedDataChannel                  Channel = 11  // Speed Data
	BeaconPulseChannel                Channel = 12  // Beacon Pulse
	FrequencyChannel1                 Channel = 14  // Frequency 1
	FrequencyChannel2                 Channel = 15  // Frequency 2
	FrequencyChannel3                 Channel = 16  // Frequency 3
	FrequencyChannel4                 Channel = 17  // Frequency 4
	FrequencyChannel5                 Channel = 18  // Frequency 5
	AnalogChannel1                    Channel = 20  // Analog 1
	AnalogChannel2                    Channel = 21  // Analog 2
	AnalogChannel3                    Channel = 22  // Analog 3
	AnalogChannel4                    Channel = 23  // Analog 4
	AnalogChannel5                    Channel = 24  // Analog 5
	AnalogChannel6                    Channel = 25  // Analog 6
	AnalogChannel7                    Channel = 26  // Analog 7
	AnalogChannel8                    Channel = 27  // Analog 8
	AnalogChannel9                    Channel = 28  // Analog 9
	AnalogChannel10                   Channel = 29  // Analog 10
	AnalogChannel11                   Channel = 30  // Analog 11
	AnalogChannel12                   Channel = 31  // Analog 12
	AnalogChannel13                   Channel = 32  // Analog 13
	AnalogChannel14                   Channel = 33  // Analog 14
	AnalogChannel15                   Channel = 34  // Analog 15
	AnalogChannel16                   Channel = 35  // Analog 16
	AnalogChannel17                   Channel = 36  // Analog 17
	AnalogChannel18                   Channel = 37  // Analog 18
	AnalogChannel19                   Channel = 38  // Analog 19
	AnalogChannel20                   Channel = 39  // Analog 20
	AnalogChannel21                   Channel = 40  // Analog 21
	AnalogChannel22                   Channel = 41  // Analog 22
	AnalogChannel23                   Channel = 42  // Analog 23
	AnalogChannel24                   Channel = 43  // Analog 24
	AnalogChannel25                   Channel = 44  // Analog 25
	AnalogChannel26                   Channel = 45  // Analog 26
	AnalogChannel27                   Channel = 46  // Analog 27
	AnalogChannel28                   Channel = 47  // Analog 28
	AnalogChannel29                   Channel = 48  // Analog 29
	AnalogChannel30                   Channel = 49  // Analog 30
	AnalogChannel31                   Channel = 50  // Analog 31
	AnalogChannel32                   Channel = 51  // Analog 32
	ChannelDataChannel                Channel = 52  // Channel Data
	DisplayDataChannel                Channel = 53  // Display Data
	ReflashChannel                    Channel = 54  // Reflash
	DateStorageChannel                Channel = 55  // Data Storage
	CourseDataChannel                 Channel = 56  // Course Data
	GPSAltitudeChannel                Channel = 57  // GPS Altitude
	ExtendedFrequencyChannel1         Channel = 58  // Extended Frequency 1
	ExtendedFrequencyChannel2         Channel = 59  // Extended Frequency 2
	ExtendedFrequencyChannel3         Channel = 60  // Extended Frequency 3
	ExtendedFrequencyChannel4         Channel = 61  // Extended Frequency 4
	ExtendedRPMChannel                Channel = 62  // Extended RPM
	StartRunChannel                   Channel = 63  // Start Run
	ProcessedSpeedDataChannel         Channel = 64  // Processed Speed Data
	GearSetupDataChannel              Channel = 65  // Gear Setup Data
	BargraphSetupDataChannel          Channel = 66  // Bargraph Setup Data
	DashboardSetupDataChannel         Channel = 67  // Dashboard Setup Data 1
	DashboardSetupDataTwoChannel      Channel = 68  // Dashboard Setup Data 2
	NewTargetSectorTimeChannel        Channel = 69  // New Target Sector Time
	NewTargetMarkerTimeChannel        Channel = 70  // New Target Marker Time
	AuxilleryInputModuleNumberChannel Channel = 71  // Auxillery Input Module Number
	ExternalTemperatureChannel        Channel = 72  // External Temperature
	ExternalFrequencyChannel          Channel = 73  // External Frequency
	ExternalPercentageChannel         Channel = 74  // External Percentage
	ExternalTimeChannel               Channel = 75  // External Time
	NewLCDDataChannel                 Channel = 76  // New LCD Data
	NewLEDDataChannel                 Channel = 77  // New LED Data
	PreCalculatedDistanceDataChannel  Channel = 78  // Pre-Calculated Distance Data
	YawRatesChannel                   Channel = 79  // Yaw Rates
	CalculatedYawChannel              Channel = 80  // Calculated Yaw
	PitchRateChannel                  Channel = 81  // Pitch Rate
	PitchAngleChannel                 Channel = 82  // Pitch Angle
	RollRateChannel                   Channel = 83  // Roll Rate
	RollAngleChannel                  Channel = 84  // Roll Angle
	GradientChannel                   Channel = 85  // Gradient
	PulseCountChannel1                Channel = 86  // Pulse Count 1
	PulseCountChannel2                Channel = 87  // Pulse Count 2
	PulseCountChannel3                Channel = 88  // Pulse Count 3
	PulseCountChannel4                Channel = 89  // Pulse Count 4
	BaselineChannel                   Channel = 90  // Baseline
	UnitControlChannel                Channel = 91  // Unit Control
	ZAccelerationChannel              Channel = 92  // Z Acceleration
	ExternalAngleChannel              Channel = 93  // External Angle
	ExternalPressureChannel           Channel = 94  // External Pressure
	ExternalMiscellaneousChannel      Channel = 95  // External Misc
	SectorDefinitionChannel           Channel = 101 // Sector Definition
	DVRCommunicationChannel           Channel = 103 // DVC Communications
	VideoFrameIndexChannel            Channel = 104 // Video Frame Index
)

var channelLengths = map[Channel]int{
	RunInformationChannel:             9,
	RunStatusChannel:                  11,
	NewSectorTime:                     7,
	LapMarkerChannel:                  21,
	LoggerStorageChannel:              6,
	GPSTimeStorageChannel:             6,
	AccelerationsChannel:              6,
	TimestampChannel:                  5,
	GPSPositionChannel:                14,
	SpeedDataChannel:                  10,
	BeaconPulseChannel:                3,
	FrequencyChannel1:                 5,
	FrequencyChannel2:                 5,
	FrequencyChannel3:                 5,
	FrequencyChannel4:                 5,
	FrequencyChannel5:                 5,
	AnalogChannel1:                    4,
	AnalogChannel2:                    4,
	AnalogChannel3:                    4,
	AnalogChannel4:                    4,
	AnalogChannel5:                    4,
	AnalogChannel6:                    4,
	AnalogChannel7:                    4,
	AnalogChannel8:                    4,
	AnalogChannel9:                    4,
	AnalogChannel10:                   4,
	AnalogChannel11:                   4,
	AnalogChannel12:                   4,
	AnalogChannel13:                   4,
	AnalogChannel14:                   4,
	AnalogChannel15:                   4,
	AnalogChannel16:                   4,
	AnalogChannel17:                   4,
	AnalogChannel18:                   4,
	AnalogChannel19:                   4,
	AnalogChannel20:                   4,
	AnalogChannel21:                   4,
	AnalogChannel22:                   4,
	AnalogChannel23:                   4,
	AnalogChannel24:                   4,
	AnalogChannel25:                   4,
	AnalogChannel26:                   4,
	AnalogChannel27:                   4,
	AnalogChannel28:                   4,
	AnalogChannel29:                   4,
	AnalogChannel30:                   4,
	AnalogChannel31:                   4,
	AnalogChannel32:                   4,
	ChannelDataChannel:                67,
	DisplayDataChannel:                11,
	ReflashChannel:                    6,
	DateStorageChannel:                10,
	CourseDataChannel:                 10,
	GPSAltitudeChannel:                10,
	ExtendedFrequencyChannel1:         11,
	ExtendedFrequencyChannel2:         11,
	ExtendedFrequencyChannel3:         11,
	ExtendedFrequencyChannel4:         11,
	ExtendedRPMChannel:                11,
	StartRunChannel:                   3,
	ProcessedSpeedDataChannel:         5,
	GearSetupDataChannel:              30,
	BargraphSetupDataChannel:          11,
	DashboardSetupDataChannel:         4,
	DashboardSetupDataTwoChannel:      4,
	NewTargetSectorTimeChannel:        42,
	NewTargetMarkerTimeChannel:        42,
	AuxilleryInputModuleNumberChannel: 3,
	ExternalTemperatureChannel:        5,
	ExternalFrequencyChannel:          5,
	ExternalPercentageChannel:         5,
	ExternalTimeChannel:               6,
	NewLCDDataChannel:                 24,
	NewLEDDataChannel:                 3,
	PreCalculatedDistanceDataChannel:  6,
	YawRatesChannel:                   4,
	CalculatedYawChannel:              4,
	PitchRateChannel:                  5,
	PitchAngleChannel:                 5,
	RollRateChannel:                   5,
	RollAngleChannel:                  5,
	GradientChannel:                   10,
	PulseCountChannel1:                5,
	PulseCountChannel2:                5,
	PulseCountChannel3:                5,
	PulseCountChannel4:                5,
	BaselineChannel:                   6,
	UnitControlChannel:                5,
	ZAccelerationChannel:              4,
	ExternalAngleChannel:              5,
	ExternalPressureChannel:           6,
	ExternalMiscellaneousChannel:      5,
	SectorDefinitionChannel:           19,
	DVRCommunicationChannel:           18,
	VideoFrameIndexChannel:            6,
}
