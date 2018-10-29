package model

import "encoding/xml"

type Gtrxgroup struct {
	XMLName xml.Name `xml:"GTRXGROUP"`
	ATTRIBUTES GtrxgroupAttributes `xml:"attributes"`
}

type GtrxgroupAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	GTRXGROUPID string `xml:"GTRXGROUPID"`
	GLOCELLID string `xml:"GLOCELLID"`
	SNDMD string `xml:"SNDMD"`
	RCVMD string `xml:"RCVMD"`
	WORKMODE string `xml:"WORKMODE"`
	USERLABEL string `xml:"USERLABEL"`
	OBJID string `xml:"OBJID"`
}

