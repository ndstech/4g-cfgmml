package model

import "encoding/xml"

type Gbtsabiscp struct {
	XMLName xml.Name `xml:"GBTSABISCP"`
	ATTRIBUTES GbtsabiscpAttributes `xml:"attributes"`
}

type GbtsabiscpAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	ABISCPID string `xml:"ABISCPID"`
	CPBEARID string `xml:"CPBEARID"`
	OBJID string `xml:"OBJID"`
}

