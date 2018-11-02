package model

import "encoding/xml"

type Almport struct {
	XMLName    xml.Name          `xml:"ALMPORT"`
	ATTRIBUTES AlmportAttributes `xml:"attributes"`
}

type AlmportAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN      string   `xml:"CN"`
	SRN     string   `xml:"SRN"`
	SN      string   `xml:"SN"`
	PN      string   `xml:"PN"`
	SW      string   `xml:"SW"`
	AID     string   `xml:"AID"`
	PT      string   `xml:"PT"`
	AVOL    string   `xml:"AVOL"`
	DTPRD   string   `xml:"DTPRD"`
	UL      string   `xml:"UL"`
	LL      string   `xml:"LL"`
	ST      string   `xml:"ST"`
	SMUL    string   `xml:"SMUL"`
	SMLL    string   `xml:"SMLL"`
	SOUL    string   `xml:"SOUL"`
	SOLL    string   `xml:"SOLL"`
}
