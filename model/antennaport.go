package model

import "encoding/xml"

type Antennaport struct {
	XMLName xml.Name `xml:"ANTENNAPORT"`
	ATTRIBUTES AntennaportAttributes `xml:"attributes"`
}

type AntennaportAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	PN string `xml:"PN"`
	FEEDERLENGTH string `xml:"FEEDERLENGTH"`
	DLDELAY string `xml:"DLDELAY"`
	ULDELAY string `xml:"ULDELAY"`
	PWRSWITCH string `xml:"PWRSWITCH"`
	THRESHOLDTYPE string `xml:"THRESHOLDTYPE"`
	UOTHD string `xml:"UOTHD"`
	UCTHD string `xml:"UCTHD"`
	OOTHD string `xml:"OOTHD"`
	OCTHD string `xml:"OCTHD"`
	ULTRADELAYSW string `xml:"ULTRADELAYSW"`
}

