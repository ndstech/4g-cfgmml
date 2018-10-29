package model

import "encoding/xml"

type Emu struct {
	XMLName xml.Name `xml:"EMU"`
	ATTRIBUTES EmuAttributes `xml:"attributes"`
}

type EmuAttributes struct {
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
	HLTHD string `xml:"HLTHD"`
	HUTHD string `xml:"HUTHD"`
	SAAF string `xml:"SAAF"`
	SBAF string `xml:"SBAF"`
}

