package osm

import (
	"encoding/xml"
	"io"
	"os"
	"strings"
	"time"
)

// Osm struct
type Map struct {
	Bounds    Bounds
	Nodes     []Node
	Ways      []Way
	Relations []Relation
}

// Bounds struct
type Bounds struct {
	XMLName xml.Name `xml:"bounds"`
	Minlat  float64  `xml:"minlat,attr"`
	Minlon  float64  `xml:"minlon,attr"`
	Maxlat  float64  `xml:"maxlat,attr"`
	Maxlon  float64  `xml:"maxlon,attr"`
}

// Location struct
type Location struct {
	Type        string
	Coordinates []float64
}

// Tag struct
type Tag struct {
	XMLName xml.Name `xml:"tag"`
	Key     string   `xml:"k,attr"`
	Value   string   `xml:"v,attr"`
}

// Elem is a OSM base element
type Elem struct {
	ID        int64 `xml:"id,attr"`
	Loc       Location
	Version   int       `xml:"version,attr"`
	Ts        time.Time `xml:"timestamp,attr"`
	UID       int64     `xml:"uid,attr"`
	User      string    `xml:"user,attr"`
	ChangeSet int64     `xml:"changeset,attr"`
}

// Node structure
type Node struct {
	Elem
	XMLName xml.Name `xml:"node"`
	Lat     float64  `xml:"lat,attr"`
	Lng     float64  `xml:"lon,attr"`
	Tag     []Tag    `xml:"tag"`
}

// Way struct
type Way struct {
	Elem
	XMLName xml.Name `xml:"way"`
	Tags    map[string]string
	RTags   []Tag `xml:"tag"`
	Nds     []struct {
		ID int64 `xml:"ref,attr"`
	} `xml:"nd"`
}

// Member struct
type Member struct {
	Type string `xml:"type,attr"`
	Ref  int64  `xml:"ref,attr"`
	Role string `xml:"role,attr"`
}

// Relation struct
type Relation struct {
	Elem
	Visible bool     `xml:"visible,attr"`
	Version string   `xml:"version,attr"`
	Members []Member `xml:"member"`
	Tags    []Tag    `xml:"tag"`
}

// DecodeFile an OSM file
func DecodeFile(fileName string) (*Map, error) {

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return Decode(file)
}

func DecodeString(data string) (*Map, error) {
	return Decode(strings.NewReader(data))
}

// Decode an reader
func Decode(reader io.Reader) (*Map, error) {
	var (
		o   = new(Map)
		err error
	)

	decoder := xml.NewDecoder(reader)
	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}

		switch typedToken := token.(type) {
		case xml.StartElement:
			switch typedToken.Name.Local {
			case "bounds":
				var b Bounds
				err = decoder.DecodeElement(&b, &typedToken)
				if err != nil {
					return nil, err
				}
				o.Bounds = b

			case "node":
				var n Node
				err = decoder.DecodeElement(&n, &typedToken)
				if err != nil {
					return nil, err
				}
				o.Nodes = append(o.Nodes, n)

			case "way":
				var w Way
				err = decoder.DecodeElement(&w, &typedToken)
				if err != nil {
					return nil, err
				}
				o.Ways = append(o.Ways, w)

			case "relation":
				var r Relation
				err = decoder.DecodeElement(&r, &typedToken)
				if err != nil {
					return nil, err
				}
				o.Relations = append(o.Relations, r)
			}
		}
	}
	return o, nil
}
