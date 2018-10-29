package model

import "encoding/xml"

type Lineclk struct {
	XMLName xml.Name `xml:"LINECLK"`
	ATTRIBUTES LineclkAttributes `xml:"attributes"`
}

type LineclkAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LN string `xml:"LN"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	PT string `xml:"PT"`
	PN string `xml:"PN"`
	PRI string `xml:"PRI"`
}

