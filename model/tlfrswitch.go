package model

import "encoding/xml"

type Tlfrswitch struct {
	XMLName xml.Name `xml:"TLFRSWITCH"`
	ATTRIBUTES TlfrswitchAttributes `xml:"attributes"`
}

type TlfrswitchAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	TLFRSWITCH string `xml:"TLFRSWITCH"`
}

