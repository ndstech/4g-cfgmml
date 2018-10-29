package model

import "encoding/xml"

type Psuis struct {
	XMLName xml.Name `xml:"PSUIS"`
	ATTRIBUTES PsuisAttributes `xml:"attributes"`
}

type PsuisAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	PSUISS string `xml:"PSUISS"`
}

