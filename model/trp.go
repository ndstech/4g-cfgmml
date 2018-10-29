package model

import "encoding/xml"

type Trp struct {
	XMLName xml.Name `xml:"TRP"`
	ATTRIBUTES TrpAttributes `xml:"attributes"`
}

type TrpAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	SUBTYPE string `xml:"SUBTYPE"`
}

