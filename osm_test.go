package osm

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestUnmarshalFile(t *testing.T) {
	testFileName := "testfile.osm"
	err := ioutil.WriteFile(testFileName, testData, 0600)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		os.Remove(testFileName)
	}()

	o, err := DecodeFile(testFileName)
	if err != nil {
		t.Fatal(err)
	}
	if len(o.Nodes) != 1 {
		t.Fatal("Map has only one node")
	}
	if len(o.Ways) != 1 {
		t.Fatal("Map has only one way")
	}
	if len(o.Relations) != 1 {
		t.Fatal("Map has only one relation")
	}
}

func TestUnmarshalData(t *testing.T) {
	o, e := Decode(bytes.NewReader(testData))
	if e != nil {
		t.Fatal(e)
	}
	if len(o.Nodes) != 1 {
		t.Fatal("Map has only one node")
	}
	if len(o.Ways) != 1 {
		t.Fatal("Map has only one way")
	}
	if len(o.Relations) != 1 {
		t.Fatal("Map has only one relation")
	}
}

var testData = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<osm version="0.6" generator="CGImap 0.5.8 (16052 thorn-02.openstreetmap.org)" copyright="OpenStreetMap and contributors" attribution="http://www.openstreetmap.org/copyright" license="http://opendatacommons.org/licenses/odbl/1-0/">
 <bounds minlat="64.0918000" minlon="-21.9304200" maxlat="64.0928200" maxlon="-21.9262200"/>
 <node id="14586443" visible="true" version="3" changeset="7510701" timestamp="2011-03-10T04:15:27Z" user="Kjarrval" uid="209717" lat="64.0912791" lon="-21.9271369"/>
 <way id="23341403" visible="true" version="15" changeset="43895070" timestamp="2016-11-23T12:34:38Z" user="Baldvin" uid="50061">
  <nd ref="137868465"/>
  <nd ref="1194972474"/>
  <nd ref="4517020462"/>
  <nd ref="1196380946"/>
  <nd ref="1196381066"/>
  <nd ref="2470653354"/>
  <nd ref="2313402455"/>
  <nd ref="2313402451"/>
  <nd ref="252745252"/>
  <nd ref="2313403164"/>
  <nd ref="1196381130"/>
  <nd ref="1196381560"/>
  <nd ref="1196381263"/>
  <nd ref="1689245704"/>
  <tag k="highway" v="tertiary"/>
  <tag k="name" v="Vífilsstaðavegur"/>
 </way>
 <relation id="65930" visible="true" version="53" changeset="45713748" timestamp="2017-02-01T11:51:31Z" user="mrpulley" uid="67896">
  <member type="way" ref="22560575" role="forward"/>
  <tag k="name" v="Hafnarfjarðarvegur"/>
  <tag k="network" v="S"/>
  <tag k="ref" v="40"/>
  <tag k="route" v="road"/>
  <tag k="type" v="route"/>
 </relation>
</osm>`)
