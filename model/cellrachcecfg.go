package model

import "encoding/xml"

type Cellrachcecfg struct {
	XMLName xml.Name `xml:"CELLRACHCECFG"`
	ATTRIBUTES CellrachcecfgAttributes `xml:"attributes"`
}

type CellrachcecfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	CoverageLevel string `xml:"CoverageLevel"`
	ContentionResolutionTimer string `xml:"ContentionResolutionTimer"`
	PrachTransmissionPeriod string `xml:"PrachTransmissionPeriod"`
	PrachSubcarrierOffset string `xml:"PrachSubcarrierOffset"`
	PrachRepetitionCount string `xml:"PrachRepetitionCount"`
	MaxNumPreambleAttempt string `xml:"MaxNumPreambleAttempt"`
	PrachDetectionThld string `xml:"PrachDetectionThld"`
	PrachStartTime string `xml:"PrachStartTime"`
}

