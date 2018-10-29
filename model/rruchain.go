package model

import "encoding/xml"

type Rruchain struct {
	XMLName xml.Name `xml:"RRUCHAIN"`
	ATTRIBUTES RruchainAttributes `xml:"attributes"`
}

type RruchainAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	RCN string `xml:"RCN"`
	TT string `xml:"TT"`
	BM string `xml:"BM"`
	HCN string `xml:"HCN"`
	HSRN string `xml:"HSRN"`
	HSN string `xml:"HSN"`
	HPN string `xml:"HPN"`
	BRKPOS1 string `xml:"BRKPOS1"`
	BRKPOS2 string `xml:"BRKPOS2"`
	AT string `xml:"AT"`
	CR string `xml:"CR"`
	LSN string `xml:"LSN"`
	PROTOCOL string `xml:"PROTOCOL"`
	SBT string `xml:"SBT"`
	CONNPORTNUM string `xml:"CONNPORTNUM"`
}

