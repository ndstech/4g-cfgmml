package model

import "encoding/xml"

type Peu struct {
	XMLName xml.Name `xml:"PEU"`
	ATTRIBUTES PeuAttributes `xml:"attributes"`
}

type PeuAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
}

