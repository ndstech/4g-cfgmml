package model

import "encoding/xml"

type Tddframeoffset struct {
	XMLName xml.Name `xml:"TDDFRAMEOFFSET"`
	ATTRIBUTES TddframeoffsetAttributes `xml:"attributes"`
}

type TddframeoffsetAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	TimeOffset string `xml:"TimeOffset"`
	ObjId string `xml:"objId"`
}

