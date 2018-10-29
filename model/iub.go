package model

import "encoding/xml"

type Iub struct {
	XMLName xml.Name `xml:"IUB"`
	ATTRIBUTES IubAttributes `xml:"attributes"`
}

type IubAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	IUBID string `xml:"IUBID"`
	OBJID string `xml:"OBJID"`
	UPEPGROUPID string `xml:"UPEPGROUPID"`
	USERLABEL string `xml:"USERLABEL"`
	STATICCHKSW string `xml:"STATICCHKSW"`
}

