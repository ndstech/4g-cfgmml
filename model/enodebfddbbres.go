package model

import "encoding/xml"

type Enodebfddbbres struct {
	XMLName xml.Name `xml:"eNodeBFddBbRes"`
	ATTRIBUTES EnodebfddbbresAttributes `xml:"attributes"`
}

type EnodebfddbbresAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	ENodeBFddBbResId string `xml:"eNodeBFddBbResId"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	Capacity string `xml:"Capacity"`
	AdmState string `xml:"AdmState"`
	ObjId string `xml:"objId"`
}

