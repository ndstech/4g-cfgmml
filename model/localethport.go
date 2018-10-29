package model

import "encoding/xml"

type Localethport struct {
	XMLName xml.Name `xml:"LOCALETHPORT"`
	ATTRIBUTES LocalethportAttributes `xml:"attributes"`
}

type LocalethportAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	SWITCH string `xml:"SWITCH"`
	GRATUITOUSARPSW string `xml:"GRATUITOUSARPSW"`
	IP6SW string `xml:"IP6SW"`
}

