package model

import "encoding/xml"

type Retport struct {
	XMLName xml.Name `xml:"RETPORT"`
	ATTRIBUTES RetportAttributes `xml:"attributes"`
}

type RetportAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	PN string `xml:"PN"`
	PWRSWITCH string `xml:"PWRSWITCH"`
	THRESHOLDTYPE string `xml:"THRESHOLDTYPE"`
	UOTHD string `xml:"UOTHD"`
	UCTHD string `xml:"UCTHD"`
	OOTHD string `xml:"OOTHD"`
	OCTHD string `xml:"OCTHD"`
}

