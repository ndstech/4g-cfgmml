package model

import "encoding/xml"

type Cellmbfcs struct {
	XMLName xml.Name `xml:"CELLMBFCS"`
	ATTRIBUTES CellmbfcsAttributes `xml:"attributes"`
}

type CellmbfcsAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	CellSyntheticalRateHoHyst string `xml:"CellSyntheticalRateHoHyst"`
	CellTrafficLoadThd string `xml:"CellTrafficLoadThd"`
	InterFreqMeasLoadRatioThd string `xml:"InterFreqMeasLoadRatioThd"`
	InterFreqMeasMcsThd string `xml:"InterFreqMeasMcsThd"`
	InterFreqMeasServiceThd string `xml:"InterFreqMeasServiceThd"`
	MeasInfoUpdInterval string `xml:"MeasInfoUpdInterval"`
	UeSpectralEffHoHyst string `xml:"UeSpectralEffHoHyst"`
	LoadDiffRatioThld string `xml:"LoadDiffRatioThld"`
}

