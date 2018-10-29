package model

import "encoding/xml"

type Lioptrule struct {
	XMLName xml.Name `xml:"LIOPTRULE"`
	ATTRIBUTES LioptruleAttributes `xml:"attributes"`
}

type LioptruleAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	RuleID string `xml:"RuleID"`
	ActionType string `xml:"ActionType"`
	AtomRuleRelationType string `xml:"AtomRuleRelationType"`
	Period string `xml:"Period"`
	Action string `xml:"Action"`
	IOptFunctionID string `xml:"IOptFunctionID"`
	PenaltyTime string `xml:"PenaltyTime"`
	AdaptiveRAT string `xml:"AdaptiveRAT"`
	ActiveStatus string `xml:"ActiveStatus"`
	ObjId string `xml:"objId"`
}

