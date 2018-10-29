package model

import "encoding/xml"

type Tcu struct {
	XMLName xml.Name `xml:"TCU"`
	ATTRIBUTES TcuAttributes `xml:"attributes"`
}

type TcuAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	MCN string `xml:"MCN"`
	MSRN string `xml:"MSRN"`
	MPN string `xml:"MPN"`
	ADDR string `xml:"ADDR"`
	TLTHD string `xml:"TLTHD"`
	TUTHD string `xml:"TUTHD"`
	SBAF string `xml:"SBAF"`
	TCMODE string `xml:"TCMODE"`
}

