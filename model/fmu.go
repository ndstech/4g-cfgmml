package model

import "encoding/xml"

type Fmu struct {
	XMLName xml.Name `xml:"FMU"`
	ATTRIBUTES FmuAttributes `xml:"attributes"`
}

type FmuAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	MCN string `xml:"MCN"`
	MSRN string `xml:"MSRN"`
	MPN string `xml:"MPN"`
	ADDR string `xml:"ADDR"`
	SBAF string `xml:"SBAF"`
	STC string `xml:"STC"`
	TCMODE string `xml:"TCMODE"`
}

