package model

import "encoding/xml"

type Typdrbbsr struct {
	XMLName xml.Name `xml:"TypDrbBsr"`
	ATTRIBUTES TypdrbbsrAttributes `xml:"attributes"`
}

type TypdrbbsrAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	Qci string `xml:"Qci"`
	TPerodicBSRTimer string `xml:"TPerodicBSRTimer"`
	RetxBsrTimer string `xml:"RetxBsrTimer"`
	ObjId string `xml:"objId"`
}

