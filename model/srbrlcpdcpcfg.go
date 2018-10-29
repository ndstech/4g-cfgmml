package model

import "encoding/xml"

type Srbrlcpdcpcfg struct {
	XMLName xml.Name `xml:"SRBRLCPDCPCFG"`
	ATTRIBUTES SrbrlcpdcpcfgAttributes `xml:"attributes"`
}

type SrbrlcpdcpcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	SrbId string `xml:"SrbId"`
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
	ObjId string `xml:"objId"`
	NbENodeBPollRetransTimer string `xml:"NbENodeBPollRetransTimer"`
	NbUePollRetransTimer string `xml:"NbUePollRetransTimer"`
}

