package model

import "encoding/xml"

type Rlcpdcpparagroup struct {
	XMLName xml.Name `xml:"RLCPDCPPARAGROUP"`
	ATTRIBUTES RlcpdcpparagroupAttributes `xml:"attributes"`
}

type RlcpdcpparagroupAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	RlcPdcpParaGroupId string `xml:"RlcPdcpParaGroupId"`
	DiscardTimer string `xml:"DiscardTimer"`
	RlcMode string `xml:"RlcMode"`
	PdcpSnSize string `xml:"PdcpSnSize"`
	UlRlcSnSize string `xml:"UlRlcSnSize"`
	DlRlcSnSize string `xml:"DlRlcSnSize"`
	UeUmReorderingTimer string `xml:"UeUmReorderingTimer"`
	ENodeBUmReorderingTimer string `xml:"ENodeBUmReorderingTimer"`
	NonsptUmUeAdaptSwitch string `xml:"NonsptUmUeAdaptSwitch"`
	UlDlDiscardtimerSwitch string `xml:"UlDlDiscardtimerSwitch"`
	ObjId string `xml:"objId"`
	UeMaxRetxThreshold string `xml:"UeMaxRetxThreshold"`
	ENodeBMaxRetxThreshold string `xml:"ENodeBMaxRetxThreshold"`
	UePollByte string `xml:"UePollByte"`
	ENodeBPollByte string `xml:"ENodeBPollByte"`
	UePollPdu string `xml:"UePollPdu"`
	ENodeBPollPdu string `xml:"ENodeBPollPdu"`
	UePollRetransmitTimer string `xml:"UePollRetransmitTimer"`
	ENodeBPollRetransmitTimer string `xml:"ENodeBPollRetransmitTimer"`
	UeStatusProhibitTimer string `xml:"UeStatusProhibitTimer"`
	ENodeBStatusProhibitTimer string `xml:"ENodeBStatusProhibitTimer"`
	UeAmReorderingTimer string `xml:"UeAmReorderingTimer"`
	ENodeBAmReorderingTimer string `xml:"ENodeBAmReorderingTimer"`
	PdcpStatusRptReq string `xml:"PdcpStatusRptReq"`
	RlcParaAdaptSwitch string `xml:"RlcParaAdaptSwitch"`
	ENodeBPollRtrTimerPreset string `xml:"eNodeBPollRtrTimerPreset"`
	ENodeBStatProhTimerPreset string `xml:"eNodeBStatProhTimerPreset"`
	UePollRtrTimerPreset string `xml:"UePollRtrTimerPreset"`
	UeStatProhTimerPreset string `xml:"UeStatProhTimerPreset"`
	CaUeRlcParaAdptiveThd string `xml:"CaUeRlcParaAdptiveThd"`
	CaUeReorderingTimer string `xml:"CaUeReorderingTimer"`
	CaUeStatProhTimer string `xml:"CaUeStatProhTimer"`
	AmPdcpSnSize string `xml:"AmPdcpSnSize"`
	ENodeBReorderingTimerAdapt string `xml:"ENodeBReorderingTimerAdapt"`
	UePdcpReorderingTimer string `xml:"UePdcpReorderingTimer"`
	CatType string `xml:"CatType"`
	NbPdcpDiscardTimer string `xml:"NbPdcpDiscardTimer"`
	NbDlPdcpDiscardTimer string `xml:"NbDlPdcpDiscardTimer"`
	NbEnodebPollRetxTimer string `xml:"NbEnodebPollRetxTimer"`
	NbUePollRetxTimer string `xml:"NbUePollRetxTimer"`
}

