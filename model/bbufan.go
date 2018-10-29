package model

import "encoding/xml"

type Bbufan struct {
	XMLName xml.Name `xml:"BBUFAN"`
	ATTRIBUTES BbufanAttributes `xml:"attributes"`
}

type BbufanAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
}

