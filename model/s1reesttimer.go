package model

import "encoding/xml"

type S1reesttimer struct {
	XMLName xml.Name `xml:"S1ReestTimer"`
	ATTRIBUTES S1reesttimerAttributes `xml:"attributes"`
}

type S1reesttimerAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	S1ReestMinTimer string `xml:"S1ReestMinTimer"`
	S1ReestMaxTimer string `xml:"S1ReestMaxTimer"`
	ObjId string `xml:"objId"`
}

