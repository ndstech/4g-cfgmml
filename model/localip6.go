package model

import "encoding/xml"

type Localip6 struct {
	XMLName xml.Name `xml:"LOCALIP6"`
	ATTRIBUTES Localip6Attributes `xml:"attributes"`
}

type Localip6Attributes struct {
	XMLName xml.Name `xml:"attributes"`
	IP6 string `xml:"IP6"`
	PFXLEN string `xml:"PFXLEN"`
}

