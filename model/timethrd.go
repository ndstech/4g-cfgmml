package model

import "encoding/xml"

type Timethrd struct {
	XMLName xml.Name `xml:"TIMETHRD"`
	ATTRIBUTES TimethrdAttributes `xml:"attributes"`
}

type TimethrdAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	THRD string `xml:"THRD"`
	SWITCH string `xml:"SWITCH"`
}

