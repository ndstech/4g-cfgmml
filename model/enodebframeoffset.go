package model

import "encoding/xml"

type Enodebframeoffset struct {
	XMLName xml.Name `xml:"ENODEBFRAMEOFFSET"`
	ATTRIBUTES EnodebframeoffsetAttributes `xml:"attributes"`
}

type EnodebframeoffsetAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	TddFrameOffset string `xml:"TddFrameOffset"`
	FddFrameOffset string `xml:"FddFrameOffset"`
	ObjId string `xml:"objId"`
}

