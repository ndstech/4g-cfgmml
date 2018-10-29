package model

import "encoding/xml"

type Lanswitchport struct {
	XMLName xml.Name `xml:"LANSWITCHPORT"`
	ATTRIBUTES LanswitchportAttributes `xml:"attributes"`
}

type LanswitchportAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	PORTIDX string `xml:"PORTIDX"`
	PORTTYPE string `xml:"PORTTYPE"`
	USERLABLE string `xml:"USERLABLE"`
}

