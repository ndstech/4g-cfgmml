package model

import "encoding/xml"

type Lioptatomrule struct {
	XMLName xml.Name `xml:"LIOPTATOMRULE"`
	ATTRIBUTES LioptatomruleAttributes `xml:"attributes"`
}

type LioptatomruleAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	AtomRuleID string `xml:"AtomRuleID"`
	MeasureObjType string `xml:"MeasureObjType"`
	ConditionType string `xml:"ConditionType"`
	ThresholdforNumPara string `xml:"ThresholdforNumPara"`
	MeasureObject string `xml:"MeasureObject"`
	ObjId string `xml:"objId"`
}

