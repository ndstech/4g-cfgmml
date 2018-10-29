package model

import "encoding/xml"

type Glocell struct {
	XMLName xml.Name `xml:"GLOCELL"`
	ATTRIBUTES GlocellAttributes `xml:"attributes"`
}

type GlocellAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	GLOCELLID string `xml:"GLOCELLID"`
	OBJID string `xml:"OBJID"`
	BASEBANDPOLICY string `xml:"BASEBANDPOLICY"`
	BASEBANDEQMID string `xml:"BASEBANDEQMID"`
	USERLABEL string `xml:"USERLABEL"`
	LOCELLTYPE string `xml:"LOCELLTYPE"`
	BBEQMPOLICY string `xml:"BBEQMPOLICY"`
}

