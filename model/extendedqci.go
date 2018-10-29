package model

import "encoding/xml"

type Extendedqci struct {
	XMLName xml.Name `xml:"EXTENDEDQCI"`
	ATTRIBUTES ExtendedqciAttributes `xml:"attributes"`
}

type ExtendedqciAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	ExtendedQci string `xml:"ExtendedQci"`
	ServiceType string `xml:"ServiceType"`
	UlschPriorityFactor string `xml:"UlschPriorityFactor"`
	DlschPriorityFactor string `xml:"DlschPriorityFactor"`
	UlMinGbr string `xml:"UlMinGbr"`
	DlMinGbr string `xml:"DlMinGbr"`
	PreAllocationWeight string `xml:"PreAllocationWeight"`
	InterRatPolicyCfgGroupId string `xml:"InterRatPolicyCfgGroupId"`
	PrioritisedBitRate string `xml:"PrioritisedBitRate"`
	LogicalChannelPriority string `xml:"LogicalChannelPriority"`
	RlcPdcpParaGroupId string `xml:"RlcPdcpParaGroupId"`
	FreeUserFlag string `xml:"FreeUserFlag"`
	FlowCtrlType string `xml:"FlowCtrlType"`
	LtePttDedicatedExtendedQci string `xml:"LtePttDedicatedExtendedQci"`
	LtePttDownlinkPriority string `xml:"LtePttDownlinkPriority"`
	LtePttUplinkPriority string `xml:"LtePttUplinkPriority"`
	ExtQciCounterIndex string `xml:"ExtQciCounterIndex"`
	ObjId string `xml:"objId"`
}

