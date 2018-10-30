package model

import "encoding/xml"

type Cagroup struct {
	XMLName xml.Name `xml:"CaGroup"`
	ATTRIBUTES CagroupAttributes `xml:"attributes"`
}

type CagroupAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CaGroupId string `xml:"CaGroupId"`
	CaGroupTypeInd string `xml:"CaGroupTypeInd"`
	ObjId string `xml:"objId"`
}

