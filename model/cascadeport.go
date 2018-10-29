package model

import "encoding/xml"

type Cascadeport struct {
	XMLName xml.Name `xml:"CASCADEPORT"`
	ATTRIBUTES CascadeportAttributes `xml:"attributes"`
}

type CascadeportAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	PN string `xml:"PN"`
	PT string `xml:"PT"`
	SW string `xml:"SW"`
	PM string `xml:"PM"`
}

