package dl

type WayPoint struct {
	Latitude  Coordinate
	Longitude Coordinate
}

type TrackMap struct {
	Name      string
	WayPoints []WayPoint
}

var Tracks map[string]*TrackMap

func init() {
	Tracks = make(map[string]*TrackMap)
}
