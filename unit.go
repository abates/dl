//go:generate go run internal/autogen_types.go
package dl

import (
	"fmt"
	"time"
)

func formatUnit(fs fmt.State, c rune, unit string, receiver, value interface{}) {
	switch c {
	case 'v':
		if fs.Flag('#') {
			fmt.Fprintf(fs, "%T(%v)", receiver, value)
			return
		}
		fallthrough
	case 'e', 'E', 'f', 'F', 'g', 'G':
		p, pOk := fs.Precision()
		w, wOk := fs.Width()
		switch {
		case pOk && wOk:
			fmt.Fprintf(fs, "%*.*"+string(c), w, p, value)
		case pOk:
			fmt.Fprintf(fs, "%.*"+string(c), p, value)
		case wOk:
			fmt.Fprintf(fs, "%*"+string(c), w, value)
		default:
			fmt.Fprintf(fs, "%"+string(c), value)
		}
		fmt.Fprintf(fs, " %s", unit)
	default:
		fmt.Fprintf(fs, "%%!%c(%T=%v %s", c, receiver, value, unit)
	}
}

func computeGeo(buf []byte) float64 {
	value := int64(buf[0]&0x7f)<<24 | int64(buf[1])<<16 | int64(buf[2])<<8 | int64(buf[1])
	if buf[0]&0x80 == 0x80 {
		value = value - 2147483648
	}
	return float64(value)
}

func computeAcceleration(buf []byte) Acceleration {
	value := float64(buf[0]&0x7f) + (float64(buf[1]) / 0x100)
	if buf[0]&0x80 == 0 {
		value = -value
	}
	return Acceleration(value)
}

var timezones = make(map[int]*time.Location)

func zoneLookup(offset int) (zone *time.Location) {
	var found bool
	if zone, found = timezones[offset]; !found {
		zone = time.FixedZone("GMT", offset)
		timezones[offset] = zone
	}
	return
}
