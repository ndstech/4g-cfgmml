package model

import "encoding/xml"

type Vrf struct {
	XMLName xml.Name `xml:"VRF"`
	ATTRIBUTES VrfAttributes `xml:"attributes"`
}

type VrfAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	VRFIDX string `xml:"VRFIDX"`
	USERLABEL string `xml:"USERLABEL"`
}

