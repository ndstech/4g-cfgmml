package model

import "encoding/xml"

type Localip struct {
	XMLName xml.Name `xml:"LOCALIP"`
	ATTRIBUTES LocalipAttributes `xml:"attributes"`
}

type LocalipAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	IP string `xml:"IP"`
	MASK string `xml:"MASK"`
}

