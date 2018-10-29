package model

import "encoding/xml"

type Outport struct {
	XMLName xml.Name `xml:"OUTPORT"`
	ATTRIBUTES OutportAttributes `xml:"attributes"`
}

type OutportAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	PN string `xml:"PN"`
	NAME string `xml:"NAME"`
	SW string `xml:"SW"`
}

