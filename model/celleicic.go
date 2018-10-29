package model

import "encoding/xml"

type Celleicic struct {
	XMLName xml.Name `xml:"CELLEICIC"`
	ATTRIBUTES CelleicicAttributes `xml:"attributes"`
}

type CelleicicAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	AbsPattern string `xml:"AbsPattern"`
	EicicMeasureEnable string `xml:"EicicMeasureEnable"`
	A3Offset string `xml:"A3Offset"`
	AbsAdjPeriod string `xml:"AbsAdjPeriod"`
	SampleNumTarget string `xml:"SampleNumTarget"`
	AttachCellAddThd string `xml:"AttachCellAddThd"`
	WorkMode string `xml:"WorkMode"`
	StatPeriod string `xml:"StatPeriod"`
	CreAdjPeriod string `xml:"CreAdjPeriod"`
	EicicConfigUserRatioThd string `xml:"EicicConfigUserRatioThd"`
	EicicUnConfigUserRatioThd string `xml:"EicicUnConfigUserRatioThd"`
	LargeCreMicroNumThd string `xml:"LargeCreMicroNumThd"`
	EicicAdaptiveSwitch string `xml:"EicicAdaptiveSwitch"`
	NaicNCellStatPeriod string `xml:"NaicNCellStatPeriod"`
	HighCoorCCEThd string `xml:"HighCoorCCEThd"`
	HighCoorDTXThd string `xml:"HighCoorDTXThd"`
	HighSpeedEnhancedOptSwitch string `xml:"HighSpeedEnhancedOptSwitch"`
	HSHandinUserThd string `xml:"HSHandinUserThd"`
}

