package model

import "encoding/xml"

type Timesrc struct {
	XMLName xml.Name `xml:"TIMESRC"`
	ATTRIBUTES TimesrcAttributes `xml:"attributes"`
}

type TimesrcAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	TIMESRC string `xml:"TIMESRC"`
	AUTOSWITCH string `xml:"AUTOSWITCH"`
}

