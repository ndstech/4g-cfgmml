package model

import "encoding/xml"

type Nodebfunction struct {
	XMLName xml.Name `xml:"NODEBFUNCTION"`
	ATTRIBUTES NodebfunctionAttributes `xml:"attributes"`
}

type NodebfunctionAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	OBJID string `xml:"OBJID"`
	NODEBFUNCTIONNAME string `xml:"NODEBFUNCTIONNAME"`
	APPLICATIONREF string `xml:"APPLICATIONREF"`
	NERMVERSION string `xml:"NERMVERSION"`
	USERLABEL string `xml:"USERLABEL"`
	NODEBID string `xml:"NODEBID"`
	PRODUCTVERSION string `xml:"PRODUCTVERSION"`
	INTERFACEID string `xml:"INTERFACEID"`
	LMTVERSION string `xml:"LMTVERSION"`
}

