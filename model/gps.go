package model

import "encoding/xml"

type Gps struct {
	XMLName xml.Name `xml:"GPS"`
	ATTRIBUTES GpsAttributes `xml:"attributes"`
}

type GpsAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	GN string `xml:"GN"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	WPOS string `xml:"WPOS"`
	LONG string `xml:"LONG"`
	LAT string `xml:"LAT"`
	ALT string `xml:"ALT"`
	AGL string `xml:"AGL"`
	DURATION string `xml:"DURATION"`
	PRECISION string `xml:"PRECISION"`
	CABLE_LEN string `xml:"CABLE_LEN"`
	MODE string `xml:"MODE"`
	PRI string `xml:"PRI"`
	POSCHECKSW string `xml:"POSCHECKSW"`
}

