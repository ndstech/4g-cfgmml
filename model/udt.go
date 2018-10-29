package model

import "encoding/xml"

type Udt struct {
	XMLName xml.Name `xml:"UDT"`
	ATTRIBUTES UdtAttributes `xml:"attributes"`
}

type UdtAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	UDTNO string `xml:"UDTNO"`
	UDTPARAGRPID string `xml:"UDTPARAGRPID"`
}

