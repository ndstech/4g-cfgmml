package model

import "encoding/xml"

type Qcipara struct {
	XMLName xml.Name `xml:"QciPara"`
	ATTRIBUTES QciparaAttributes `xml:"attributes"`
}

type QciparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	Qci string `xml:"Qci"`
	PreAllocationWeight string `xml:"PreAllocationWeight"`
	InterRatPolicyCfgGroupId string `xml:"InterRatPolicyCfgGroupId"`
	LogicalChannelPriority string `xml:"LogicalChannelPriority"`
	RlcPdcpParaGroupId string `xml:"RlcPdcpParaGroupId"`
	UeInactiveTimerForQci string `xml:"UeInactiveTimerForQci"`
	UlSynTimerForQci string `xml:"UlSynTimerForQci"`
	UeInactivityTimerDynDrxQci string `xml:"UeInactivityTimerDynDrxQci"`
	UeInactiveTimerPri string `xml:"UeInactiveTimerPri"`
	ObjId string `xml:"objId"`
	PrioritisedBitRate string `xml:"PrioritisedBitRate"`
	FlowCtrlType string `xml:"FlowCtrlType"`
	UlschPriorityFactor string `xml:"UlschPriorityFactor"`
	DlschPriorityFactor string `xml:"DlschPriorityFactor"`
	UlMinGbr string `xml:"UlMinGbr"`
	DlMinGbr string `xml:"DlMinGbr"`
	EmtcModeARlcParaGroupId string `xml:"EmtcModeARlcParaGroupId"`
	EmtcModeBRlcParaGroupId string `xml:"EmtcModeBRlcParaGroupId"`
	NbRlcPdcpParaGroupId string `xml:"NbRlcPdcpParaGroupId"`
	CiotUeInactiveTimer string `xml:"CiotUeInactiveTimer"`
	ServiceType string `xml:"ServiceType"`
	FreeUserFlag string `xml:"FreeUserFlag"`
	LtePttDedicatedExtendedQci string `xml:"LtePttDedicatedExtendedQci"`
	ExtQciCounterIndex string `xml:"ExtQciCounterIndex"`
}

