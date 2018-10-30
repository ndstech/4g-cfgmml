package model

import "encoding/xml"

type Cellidprdupt struct {
	XMLName xml.Name `xml:"CellIdPrdUpt"`
	ATTRIBUTES CellidprduptAttributes `xml:"attributes"`
}

type CellidprduptAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	PrdUptSwitch string `xml:"PrdUptSwitch"`
	ActionTime string `xml:"ActionTime"`
	Period string `xml:"Period"`
	ObjId string `xml:"objId"`
}

