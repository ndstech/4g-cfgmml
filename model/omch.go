package model

import "encoding/xml"

type Omch struct {
	XMLName xml.Name `xml:"OMCH"`
	ATTRIBUTES OmchAttributes `xml:"attributes"`
}

type OmchAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	FLAG string `xml:"FLAG"`
	IP string `xml:"IP"`
	MASK string `xml:"MASK"`
	PEERIP string `xml:"PEERIP"`
	PEERMASK string `xml:"PEERMASK"`
	BEAR string `xml:"BEAR"`
	BRT string `xml:"BRT"`
	CHECKTYPE string `xml:"CHECKTYPE"`
	USERLABEL string `xml:"USERLABEL"`
	RTIDX string `xml:"RTIDX"`
	BINDSECONDARYRT string `xml:"BINDSECONDARYRT"`
}

