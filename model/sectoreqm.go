package model

import "encoding/xml"

type Sectoreqm struct {
	XMLName xml.Name `xml:"SECTOREQM"`
	ATTRIBUTES SectoreqmAttributes `xml:"attributes"`
}

type SectoreqmAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	SECTOREQMID string `xml:"SECTOREQMID"`
	SECTORID string `xml:"SECTORID"`
	SECTOREQMANTENNA string `xml:"SECTOREQMANTENNA"`
	ANTCFGMODE string `xml:"ANTCFGMODE"`
	RRUCN string `xml:"RRUCN"`
	RRUSRN string `xml:"RRUSRN"`
	RRUSN string `xml:"RRUSN"`
	BEAMSHAPE string `xml:"BEAMSHAPE"`
	BEAMLAYERSPLIT string `xml:"BEAMLAYERSPLIT"`
	BEAMAZIMUTHOFFSET string `xml:"BEAMAZIMUTHOFFSET"`
}

