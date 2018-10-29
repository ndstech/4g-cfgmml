package model

import "encoding/xml"

type Vlanclass struct {
	XMLName xml.Name `xml:"VLANCLASS"`
	ATTRIBUTES VlanclassAttributes `xml:"attributes"`
}

type VlanclassAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	VLANGROUPNO string `xml:"VLANGROUPNO"`
	TRAFFIC string `xml:"TRAFFIC"`
	SRVPRIO string `xml:"SRVPRIO"`
	VLANID string `xml:"VLANID"`
	VLANPRIO string `xml:"VLANPRIO"`
}

