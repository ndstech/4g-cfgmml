package model

import "encoding/xml"

type Ueiu struct {
	XMLName xml.Name `xml:"UEIU"`
	ATTRIBUTES UeiuAttributes `xml:"attributes"`
}

type UeiuAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
}

