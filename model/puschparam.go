package model

import "encoding/xml"

type Puschparam struct {
	XMLName xml.Name `xml:"PUSCHPARAM"`
	ATTRIBUTES PuschparamAttributes `xml:"attributes"`
}

type PuschparamAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	DeltaOffsetCqiIndex string `xml:"DeltaOffsetCqiIndex"`
	DeltaOffsetRiIndex string `xml:"DeltaOffsetRiIndex"`
	DeltaOffsetAckIndex string `xml:"DeltaOffsetAckIndex"`
	ObjId string `xml:"objId"`
}

