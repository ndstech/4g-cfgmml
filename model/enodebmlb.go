package model

import "encoding/xml"

type Enodebmlb struct {
	XMLName xml.Name `xml:"eNodeBMlb"`
	ATTRIBUTES EnodebmlbAttributes `xml:"attributes"`
}

type EnodebmlbAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	InterFreqIdleMlbMode string `xml:"InterFreqIdleMlbMode"`
	InterFreqIdleMlbInterval string `xml:"InterFreqIdleMlbInterval"`
	InterFreqIdleMlbStaThd string `xml:"InterFreqIdleMlbStaThd"`
	DlCaLbMaxCCNum string `xml:"DlCaLbMaxCCNum"`
	CaUeMinDataVolRatio string `xml:"CaUeMinDataVolRatio"`
	ObjId string `xml:"objId"`
}

