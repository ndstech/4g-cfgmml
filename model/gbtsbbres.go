package model

import "encoding/xml"

type Gbtsbbres struct {
	XMLName xml.Name `xml:"GBTSBBRES"`
	ATTRIBUTES GbtsbbresAttributes `xml:"attributes"`
}

type GbtsbbresAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	GBTSBBRESID string `xml:"GBTSBBRESID"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	CAPACITY string `xml:"CAPACITY"`
	ADMSTATE string `xml:"ADMSTATE"`
	OBJID string `xml:"OBJID"`
}

