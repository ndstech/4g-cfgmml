package model

import "encoding/xml"

type E1t1bear struct {
	XMLName xml.Name `xml:"E1T1BEAR"`
	ATTRIBUTES E1t1bearAttributes `xml:"attributes"`
}

type E1t1bearAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	MODE string `xml:"MODE"`
}

