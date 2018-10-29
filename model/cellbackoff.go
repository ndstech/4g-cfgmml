package model

import "encoding/xml"

type Cellbackoff struct {
	XMLName xml.Name `xml:"CELLBACKOFF"`
	ATTRIBUTES CellbackoffAttributes `xml:"attributes"`
}

type CellbackoffAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	InterFreqBackoffPrbThd string `xml:"InterFreqBackoffPrbThd"`
	InterFreqBackoffUeNumThd string `xml:"InterFreqBackoffUeNumThd"`
	RBPriMcsSelectTrigPrbThd string `xml:"RBPriMcsSelectTrigPrbThd"`
	RBPriMcsSelectStopPrbThd string `xml:"RBPriMcsSelectStopPrbThd"`
	LowestEffDlMcsThd string `xml:"LowestEffDlMcsThd"`
	LowEffDlMcsThd string `xml:"LowEffDlMcsThd"`
	LowEffDlFactorA string `xml:"LowEffDlFactorA"`
	LowEffDlFactorK string `xml:"LowEffDlFactorK"`
	LowestEffUlMcsThd string `xml:"LowestEffUlMcsThd"`
	LowEffUlMcsThd string `xml:"LowEffUlMcsThd"`
	LowEffUlFactorA string `xml:"LowEffUlFactorA"`
	LowEffUlFactorK string `xml:"LowEffUlFactorK"`
	UlHeavyTrafficTtiProporThd string `xml:"UlHeavyTrafficTtiProporThd"`
	UlTrafficMlbUeNumThd string `xml:"UlTrafficMlbUeNumThd"`
	UlHeavyTrafficJudgePeriod string `xml:"UlHeavyTrafficJudgePeriod"`
	HoInRejectPrbThd string `xml:"HoInRejectPrbThd"`
	HoInRejectUeNumThd string `xml:"HoInRejectUeNumThd"`
	RejectedHoInCause string `xml:"RejectedHoInCause"`
	BackoffCAUserHOSw string `xml:"BackoffCAUserHOSw"`
}

