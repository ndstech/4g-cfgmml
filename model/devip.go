package model

import "encoding/xml"

type Devip struct {
	XMLName xml.Name `xml:"DEVIP"`
	ATTRIBUTES DevipAttributes `xml:"attributes"`
}

type DevipAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	SBT string `xml:"SBT"`
	PT string `xml:"PT"`
	PN string `xml:"PN"`
	VRFIDX string `xml:"VRFIDX"`
	IP string `xml:"IP"`
	MASK string `xml:"MASK"`
	USERLABEL string `xml:"USERLABEL"`
	CTRLMODE string `xml:"CTRLMODE"`
}

