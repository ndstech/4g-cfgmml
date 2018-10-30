package model

import "encoding/xml"

type Highspdadaptionpara struct {
	XMLName xml.Name `xml:"HighSpdAdaptionPara"`
	ATTRIBUTES HighspdadaptionparaAttributes `xml:"attributes"`
}

type HighspdadaptionparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	UserLowSpeedJudgeNum string `xml:"UserLowSpeedJudgeNum"`
	DopplerFactor string `xml:"DopplerFactor"`
	HighSpeedUserNumThd string `xml:"HighSpeedUserNumThd"`
	HoHisJudgeHighThd string `xml:"HoHisJudgeHighThd"`
	InterfAvoidCellNum string `xml:"InterfAvoidCellNum"`
	InterfAvoidMode string `xml:"InterfAvoidMode"`
	InterfBasedPowerOff string `xml:"InterfBasedPowerOff"`
	InterfBasedRbRatio string `xml:"InterfBasedRbRatio"`
	HighspdVectorUpdatePeriod string `xml:"HighspdVectorUpdatePeriod"`
	HighspdVectorUpdateSwitch string `xml:"HighspdVectorUpdateSwitch"`
}

