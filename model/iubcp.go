package model

import "encoding/xml"

type Iubcp struct {
	XMLName xml.Name `xml:"IUBCP"`
	ATTRIBUTES IubcpAttributes `xml:"attributes"`
}

type IubcpAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CPPT string `xml:"CPPT"`
	OBJID string `xml:"OBJID"`
	CPPN string `xml:"CPPN"`
	CPBEARID string `xml:"CPBEARID"`
	BELONG string `xml:"BELONG"`
}

