package model

import "encoding/xml"

type Manresalmrpt struct {
	XMLName xml.Name `xml:"MANRESALMRPT"`
	ATTRIBUTES ManresalmrptAttributes `xml:"attributes"`
}

type ManresalmrptAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	SWITCH string `xml:"SWITCH"`
}

