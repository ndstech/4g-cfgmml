package model

import "encoding/xml"

type Gbtspath struct {
	XMLName xml.Name `xml:"GBTSPATH"`
	ATTRIBUTES GbtspathAttributes `xml:"attributes"`
}

type GbtspathAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	PATHID string `xml:"PATHID"`
	OBJID string `xml:"OBJID"`
}

