package model

import "encoding/xml"

type Enodebtddbbres struct {
	XMLName xml.Name `xml:"eNodeBTddBbRes"`
	ATTRIBUTES EnodebtddbbresAttributes `xml:"attributes"`
}

type EnodebtddbbresAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	ENodeBTddBbResId string `xml:"eNodeBTddBbResId"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	Capacity string `xml:"Capacity"`
	AdmState string `xml:"AdmState"`
	ObjId string `xml:"objId"`
}

