package model

import "encoding/xml"

type Iratncellclassmgt struct {
	XMLName xml.Name `xml:"IRATNCELLCLASSMGT"`
	ATTRIBUTES IratncellclassmgtAttributes `xml:"attributes"`
}

type IratncellclassmgtAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	RatType string `xml:"RatType"`
	StatPeriodForNcellClass string `xml:"StatPeriodForNcellClass"`
	NCellMeasNumThd string `xml:"NCellMeasNumThd"`
	ObjId string `xml:"objId"`
}

