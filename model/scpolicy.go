package model

import "encoding/xml"

type Scpolicy struct {
	XMLName xml.Name `xml:"ScPolicy"`
	ATTRIBUTES ScpolicyAttributes `xml:"attributes"`
}

type ScpolicyAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	ScAlgoSwitch string `xml:"ScAlgoSwitch"`
	VideoInitBufTime string `xml:"VideoInitBufTime"`
	ObjId string `xml:"objId"`
}

