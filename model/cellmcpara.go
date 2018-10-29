package model

import "encoding/xml"

type Cellmcpara struct {
	XMLName xml.Name `xml:"CELLMCPARA"`
	ATTRIBUTES CellmcparaAttributes `xml:"attributes"`
}

type CellmcparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	A3Offset string `xml:"A3Offset"`
	Hysteresis string `xml:"Hysteresis"`
	TimetoTrigger string `xml:"TimetoTrigger"`
	MaxReportCells string `xml:"MaxReportCells"`
	ReportAmount string `xml:"ReportAmount"`
	ReportInterval string `xml:"ReportInterval"`
	ReportQuantity string `xml:"ReportQuantity"`
	TriggerQuantity string `xml:"TriggerQuantity"`
	A6Offset string `xml:"A6Offset"`
	IntraFreqMaxReportCells string `xml:"IntraFreqMaxReportCells"`
	IntraFreqTriggerQuantity string `xml:"IntraFreqTriggerQuantity"`
	IntraFreqReportQuantity string `xml:"IntraFreqReportQuantity"`
	InterFreqMaxReportCells string `xml:"InterFreqMaxReportCells"`
	InterFreqTriggerQuantity string `xml:"InterFreqTriggerQuantity"`
	InterFreqReportQuantity string `xml:"InterFreqReportQuantity"`
	A3MCAdaptiveSwitch string `xml:"A3MCAdaptiveSwitch"`
	AoaDetectThd string `xml:"AoaDetectThd"`
	EcidReportAmount string `xml:"EcidReportAmount"`
	EcidReportInterval string `xml:"EcidReportInterval"`
}

