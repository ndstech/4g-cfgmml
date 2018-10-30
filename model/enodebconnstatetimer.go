package model

import "encoding/xml"

type Enodebconnstatetimer struct {
	XMLName xml.Name `xml:"ENodeBConnStateTimer"`
	ATTRIBUTES EnodebconnstatetimerAttributes `xml:"attributes"`
}

type EnodebconnstatetimerAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	S1MessageWaitingTimer string `xml:"S1MessageWaitingTimer"`
	X2MessageWaitingTimer string `xml:"X2MessageWaitingTimer"`
	UuMessageWaitingTimer string `xml:"UuMessageWaitingTimer"`
	Cdma1XrttHoUuPrepareTimer string `xml:"Cdma1XrttHoUuPrepareTimer"`
	Cdma1XrttHoS1WaitingTimer string `xml:"Cdma1XrttHoS1WaitingTimer"`
	Cdma1XrttHoCompleteTimer string `xml:"Cdma1XrttHoCompleteTimer"`
	CdmaHrpdHoCompleteTimer string `xml:"CdmaHrpdHoCompleteTimer"`
	CdmaHrpdHoS1WaitingTimer string `xml:"CdmaHrpdHoS1WaitingTimer"`
	CdmaHrpdHoUuPrepareTimer string `xml:"CdmaHrpdHoUuPrepareTimer"`
	WaitRrcConnSetupCmpTimer string `xml:"WaitRrcConnSetupCmpTimer"`
	SecCmpWaitingTimer string `xml:"SecCmpWaitingTimer"`
	UpUeCapInfoWaitingTimer string `xml:"UpUeCapInfoWaitingTimer"`
	FirstForwardPacketTimer string `xml:"FirstForwardPacketTimer"`
	EndMarkerTimer string `xml:"EndMarkerTimer"`
	S1MsgWaitingTimerQci1 string `xml:"S1MsgWaitingTimerQci1"`
	X2MessageWaitingTimerQci1 string `xml:"X2MessageWaitingTimerQci1"`
	UuMessageWaitingTimerQci1 string `xml:"UuMessageWaitingTimerQci1"`
	BearerActivityThd string `xml:"BearerActivityThd"`
	ObjId string `xml:"objId"`
	RrcReestActivityThld string `xml:"RrcReestActivityThld"`
	EcidWaitingTimer string `xml:"EcidWaitingTimer"`
}

