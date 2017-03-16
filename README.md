# go-osm

OpenStreetMap XML Data parser in golang

## Install

`go get github.com/glaslos/go-osm`

## Usage 

```go
package main

import (
	"fmt"
	"github.com/glaslos/go-osm"
)

var data = `<?xml version="1.0" encoding="UTF-8"?>
<osm>
  <bounds minlat="64.0918000" minlon="-21.9304200" maxlat="64.0928200" maxlon="-21.9262200"/>
  <node id="14586443" lat="64.0912791" lon="-21.9271369"/>
  <way id="23341403">
    <nd ref="137868465"/>
    <tag k="name" v="Vífilsstaðavegur"/>
  </way>
</osm>
`

func main() {
	m, _ := osm.DecodeString(data)
	// 64.0912791 -21.9271369
	fmt.Println(m.Nodes[0].Lat, m.Nodes[0].Lng)
}
```