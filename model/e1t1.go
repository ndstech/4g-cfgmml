package model

import "encoding/xml"

type E1t1 struct {
	XMLName xml.Name `xml:"E1T1"`
	ATTRIBUTES E1t1Attributes `xml:"attributes"`
}

type E1t1Attributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	PS string `xml:"PS"`
	WORKMODE string `xml:"WORKMODE"`
	FRAME string `xml:"FRAME"`
	LNCODE string `xml:"LNCODE"`
	CLKM string `xml:"CLKM"`
	SBT string `xml:"SBT"`
	PN string `xml:"PN"`
}

