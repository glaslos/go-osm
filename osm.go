package osm

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"time"
)

// Location struct
type Location struct {
	Type        string
	Coordinates []float64
}

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

// Decode an OSM file
func Decode(fileName string) {
	nodes := []Node{}
	ways := []Way{}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Can't open OSM file: " + err.Error())
	}
	decoder := xml.NewDecoder(file)
	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}

		switch typedToken := token.(type) {
		case xml.StartElement:
			if typedToken.Name.Local == "node" {
				var n Node
				err = decoder.DecodeElement(&n, &typedToken)
				if err != nil {
					log.Fatalln(err)
				}
				nodes = append(nodes, n)
			}
			if typedToken.Name.Local == "way" {
				var w Way
				err = decoder.DecodeElement(&w, &typedToken)
				if err != nil {
					log.Fatalln(err)
				}
				ways = append(ways, w)
			}
		}
	}
	fmt.Printf("Number of nodes decoded: %d\n", len(nodes))
	fmt.Printf("Number of ways decoded: %d\n", len(ways))
	fmt.Printf("%+v", nodes[1])
}
