package model

import "encoding/xml"

type Cpbearer struct {
	XMLName xml.Name `xml:"CPBEARER"`
	ATTRIBUTES CpbearerAttributes `xml:"attributes"`
}

type CpbearerAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CPBEARID string `xml:"CPBEARID"`
	FLAG string `xml:"FLAG"`
	BEARTYPE string `xml:"BEARTYPE"`
	LINKNO string `xml:"LINKNO"`
	CTRLMODE string `xml:"CTRLMODE"`
	AUTOCFGFLAG string `xml:"AUTOCFGFLAG"`
}

