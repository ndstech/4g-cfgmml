package model

import "encoding/xml"

type Standardqci struct {
	XMLName xml.Name `xml:"STANDARDQCI"`
	ATTRIBUTES StandardqciAttributes `xml:"attributes"`
}

type StandardqciAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	Qci string `xml:"Qci"`
	PreAllocationWeight string `xml:"PreAllocationWeight"`
	InterRatPolicyCfgGroupId string `xml:"InterRatPolicyCfgGroupId"`
	RlcPdcpParaGroupId string `xml:"RlcPdcpParaGroupId"`
	LogicalChannelPriority string `xml:"LogicalChannelPriority"`
	ObjId string `xml:"objId"`
	PrioritisedBitRate string `xml:"PrioritisedBitRate"`
	FlowCtrlType string `xml:"FlowCtrlType"`
	UlschPriorityFactor string `xml:"UlschPriorityFactor"`
	DlschPriorityFactor string `xml:"DlschPriorityFactor"`
	UlMinGbr string `xml:"UlMinGbr"`
	DlMinGbr string `xml:"DlMinGbr"`
}

