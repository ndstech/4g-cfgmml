package model

import "encoding/xml"

type Cpswitch struct {
	XMLName xml.Name `xml:"CPSWITCH"`
	ATTRIBUTES CpswitchAttributes `xml:"attributes"`
}

type CpswitchAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	ES string `xml:"ES"`
}

