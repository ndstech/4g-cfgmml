package model

import "encoding/xml"

type Cellmlbho struct {
	XMLName xml.Name `xml:"CELLMLBHO"`
	ATTRIBUTES CellmlbhoAttributes `xml:"attributes"`
}

type CellmlbhoAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	IdleUeSelFreqStrategy string `xml:"IdleUeSelFreqStrategy"`
	MlbHoInProtectMode string `xml:"MlbHoInProtectMode"`
	MlbHoInProtectTimer string `xml:"MlbHoInProtectTimer"`
	InterFreqMlbHoInA1ThdRsrp string `xml:"InterFreqMlbHoInA1ThdRsrp"`
	InterFreqMlbHoInA1ThdRsrq string `xml:"InterFreqMlbHoInA1ThdRsrq"`
	InterFreqMlbHoInA2ThdRsrp string `xml:"InterFreqMlbHoInA2ThdRsrp"`
	InterFreqMlbHoInA2ThdRsrq string `xml:"InterFreqMlbHoInA2ThdRsrq"`
	InterRatMlbStrategy string `xml:"InterRatMlbStrategy"`
	MlbMatchOtherFeatureMode string `xml:"MlbMatchOtherFeatureMode"`
}

