package model

import "encoding/xml"

type Lioptrulemember struct {
	XMLName xml.Name `xml:"LIOptRuleMember"`
	ATTRIBUTES LioptrulememberAttributes `xml:"attributes"`
}

type LioptrulememberAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	RuleID string `xml:"RuleID"`
	AtomRuleID string `xml:"AtomRuleID"`
	ActiveStatus string `xml:"ActiveStatus"`
	ObjId string `xml:"objId"`
}

