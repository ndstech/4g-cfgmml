package model

import "encoding/xml"

type Cellop struct {
	XMLName xml.Name `xml:"CellOp"`
	ATTRIBUTES CellopAttributes `xml:"attributes"`
}

type CellopAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	TrackingAreaId string `xml:"TrackingAreaId"`
	CellReservedForOp string `xml:"CellReservedForOp"`
	OpUlRbUsedRatio string `xml:"OpUlRbUsedRatio"`
	OpDlRbUsedRatio string `xml:"OpDlRbUsedRatio"`
	MMECfgNum string `xml:"MMECfgNum"`
	OpUeNumRatio string `xml:"OpUeNumRatio"`
	OpPttUlRbHighThd string `xml:"OpPttUlRbHighThd"`
	OpPttUlRbLowThd string `xml:"OpPttUlRbLowThd"`
	OpPttDlRbHighThd string `xml:"OpPttDlRbHighThd"`
	OpPttDlRbLowThd string `xml:"OpPttDlRbLowThd"`
	OperatorPlmnRoundPeriod string `xml:"OperatorPlmnRoundPeriod"`
	OpResourceGroupIndex string `xml:"OpResourceGroupIndex"`
	OpNonGbrResourceRatio string `xml:"OpNonGbrResourceRatio"`
	RatFreqPriorityGroupId string `xml:"RatFreqPriorityGroupId"`
}

