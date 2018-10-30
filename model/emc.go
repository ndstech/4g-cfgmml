package model

import "encoding/xml"

type Emc struct {
	XMLName xml.Name `xml:"Emc"`
	ATTRIBUTES EmcAttributes `xml:"attributes"`
}

type EmcAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CnOperatorId string `xml:"CnOperatorId"`
	EmcEnable string `xml:"EmcEnable"`
}

