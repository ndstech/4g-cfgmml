package model

import "encoding/xml"

type Gbtsfunction struct {
	XMLName xml.Name `xml:"GBTSFUNCTION"`
	ATTRIBUTES GbtsfunctionAttributes `xml:"attributes"`
}

type GbtsfunctionAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	OBJID string `xml:"OBJID"`
	GBTSFUNCTIONNAME string `xml:"GBTSFUNCTIONNAME"`
	APPLICATIONREF string `xml:"APPLICATIONREF"`
	USERLABEL string `xml:"USERLABEL"`
	NERMVERSION string `xml:"NERMVERSION"`
	PRODUCTVERSION string `xml:"PRODUCTVERSION"`
	InterfaceId string `xml:"interfaceId"`
	LmtVersion string `xml:"lmtVersion"`
}

