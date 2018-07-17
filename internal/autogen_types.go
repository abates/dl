package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"text/template"
	"time"
)

type unitType struct {
	Name        string
	Receiver    string
	Description string
	Type        string
	Unit        string
	ShortUnit   string
}

var unitTypes = []unitType{
	{"TimeOffset", "time", "is the elapsed time", "int64", "milliseconds", "ms"},
	{"Speed", "speed", "is the vehicle speed", "float64", "meters per second", "m/s"},
	{"SpeedAccuracy", "accuracy", "is the accuracy of the speed measurement", "int", "millimeters per second", "mm/s"},
	{"Coordinate", "coordinate", "is a single point on either the lattitudinal or longitudinal axis", "float64", "degrees", "°"},
	{"GPSAccuracy", "accuracy", "is the accuracy of the GPS coordinates", "int", "millimeters", "mm"},
	{"Heading", "heading", "is the direction something is headed. In the case of course information the heading is the direction the vehicle is moving. For lap markers, the heading is direction the marker is pointing", "float64", "degrees", "°"},
	{"HeadingAccuracy", "accuracy", "is the accuracy of the GPS heading", "float64", "degrees", "°"},
	{"Acceleration", "acceleration", "is the acceleration of the vehicle in a given direction", "float64", "standard gravity", "G"},
	{"GPSTime", "gpsTime", "is the number of milliseconds since midnight between Saturday and Sunday", "uint32", "milliseconds", "ms"},
	{"Voltage", "voltage", "is a measurement sampled from the anlog inputs of the data logger", "int", "millivolts", "mV"},
	{"Frequency", "freq", "is a measurement sampled from the frequency inputs of the data logger", "float64", "hertz", "hz"},
	{"Altitude", "altitude", "is the height above sea level as measured by GPS", "int", "millimeters", "mm"},
	{"AltitudeAccuracy", "accuracy", "is the accuracy of the altitude measurement", "int", "millimeters", "mm"},
}

const typeTemplate = `
{{generateComment .Name .Description .Unit}}
type {{.Name}} {{.Type}}

// Format satisfies interface fmt.Formatter
func ({{.Receiver}} {{.Name}}) Format(f fmt.State, c rune) { formatUnit(f, c, "{{.ShortUnit}}", {{.Receiver}}, {{.Type}}({{.Receiver}})) }
`

var gopath string
var unitPkgPath string

func generateComment(name, description, unit string) string {
	return fmt.Sprintf("// %s %s. %s is measured in %s", name, description, name, unit)
}

func owner() string { return "Andrew Bates" }

func main() {
	licenseText, _ := ioutil.ReadFile("internal/license.tmpl")
	funcs := template.FuncMap{"now": time.Now, "owner": owner, "generateComment": generateComment}
	license := template.Must(template.New("license").Funcs(funcs).Parse(string(licenseText)))
	form := template.Must(template.New("format").Funcs(funcs).Parse(typeTemplate))

	f, err := os.Create("unit_types.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	buf := bytes.NewBuffer(make([]byte, 0))
	err = license.Execute(buf, struct{}{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(buf, "\npackage main\n")
	fmt.Fprintf(buf, "import \"fmt\"\n")

	for _, t := range unitTypes {
		err = form.Execute(buf, t)
		if err != nil {
			log.Fatal(err)
		}
	}

	b, err := format.Source(buf.Bytes())
	if err != nil {
		f.Write(buf.Bytes()) // This is here to debug bad format
		log.Fatalf("error formatting: %s", err)
	}

	f.Write(b)
}
