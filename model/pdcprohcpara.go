package model

import "encoding/xml"

type Pdcprohcpara struct {
	XMLName xml.Name `xml:"PdcpRohcPara"`
	ATTRIBUTES PdcprohcparaAttributes `xml:"attributes"`
}

type PdcprohcparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	RohcSwitch string `xml:"RohcSwitch"`
	HighestMode string `xml:"HighestMode"`
	Profiles string `xml:"Profiles"`
	ObjId string `xml:"objId"`
}

