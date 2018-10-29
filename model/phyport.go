package model

import "encoding/xml"

type Phyport struct {
	XMLName xml.Name `xml:"PHYPORT"`
	ATTRIBUTES PhyportAttributes `xml:"attributes"`
}

type PhyportAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	PT string `xml:"PT"`
	PN string `xml:"PN"`
	ADMINISTRATIVESTATE string `xml:"ADMINISTRATIVESTATE"`
}

