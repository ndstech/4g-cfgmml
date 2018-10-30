package model

import "encoding/xml"

type Timealignmenttimer struct {
	XMLName xml.Name `xml:"TimeAlignmentTimer"`
	ATTRIBUTES TimealignmenttimerAttributes `xml:"attributes"`
}

type TimealignmenttimerAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	TimeAlignmentTimer string `xml:"TimeAlignmentTimer"`
	TimingAdvCmdOptSwitch string `xml:"TimingAdvCmdOptSwitch"`
	TimingMeasMode string `xml:"TimingMeasMode"`
	TACmdSendPeriod string `xml:"TACmdSendPeriod"`
	TimingResOptSwitch string `xml:"TimingResOptSwitch"`
	PucchTimeOffsetCompVal string `xml:"PucchTimeOffsetCompVal"`
	HighSpeedTaCmdSendPeriod string `xml:"HighSpeedTaCmdSendPeriod"`
	TaEnhance string `xml:"TaEnhance"`
}

