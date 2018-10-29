package model

import "encoding/xml"

type Portpolicy struct {
	XMLName xml.Name `xml:"PORTPOLICY"`
	ATTRIBUTES PortpolicyAttributes `xml:"attributes"`
}

type PortpolicyAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	USBACCESSSECPOLICY string `xml:"USBACCESSSECPOLICY"`
}

