package model

import "encoding/xml"

type Singleipswitch struct {
	XMLName xml.Name `xml:"SINGLEIPSWITCH"`
	ATTRIBUTES SingleipswitchAttributes `xml:"attributes"`
}

type SingleipswitchAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	SINGLEIPSW string `xml:"SINGLEIPSW"`
}

