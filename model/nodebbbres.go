package model

import "encoding/xml"

type Nodebbbres struct {
	XMLName xml.Name `xml:"NODEBBBRES"`
	ATTRIBUTES NodebbbresAttributes `xml:"attributes"`
}

type NodebbbresAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	OBJID string `xml:"OBJID"`
	NODEBBBRESID string `xml:"NODEBBBRESID"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	CAPACITY string `xml:"CAPACITY"`
	ADMSTATE string `xml:"ADMSTATE"`
	HCE string `xml:"HCE"`
}

