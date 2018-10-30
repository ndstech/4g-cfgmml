package model

import "encoding/xml"

type Cellqcipara struct {
	XMLName xml.Name `xml:"CellQciPara"`
	ATTRIBUTES CellqciparaAttributes `xml:"attributes"`
}

type CellqciparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	Qci string `xml:"Qci"`
	InterFreqHoGroupId string `xml:"InterFreqHoGroupId"`
	InterRatHoCdma1xRttGroupId string `xml:"InterRatHoCdma1xRttGroupId"`
	InterRatHoCdmaHrpdGroupId string `xml:"InterRatHoCdmaHrpdGroupId"`
	InterRatHoCommGroupId string `xml:"InterRatHoCommGroupId"`
	InterRatHoGeranGroupId string `xml:"InterRatHoGeranGroupId"`
	InterRatHoUtranGroupId string `xml:"InterRatHoUtranGroupId"`
	IntraFreqHoGroupId string `xml:"IntraFreqHoGroupId"`
	DrxParaGroupId string `xml:"DrxParaGroupId"`
	SriPeriod string `xml:"SriPeriod"`
	RlfTimerConstCfgInd string `xml:"RlfTimerConstCfgInd"`
	TrafficRelDelay string `xml:"TrafficRelDelay"`
	QciPriorityForHo string `xml:"QciPriorityForHo"`
	PreallocationParaGroupId string `xml:"PreallocationParaGroupId"`
	QciPriorityForDrx string `xml:"QciPriorityForDrx"`
	QciAlgoSwitch string `xml:"QciAlgoSwitch"`
	QciEutranFreqRelationId string `xml:"QciEutranFreqRelationId"`
	QciUtranFreqRelationId string `xml:"QciUtranFreqRelationId"`
	QciGeranFreqRelationId string `xml:"QciGeranFreqRelationId"`
	MlbQciGroupId string `xml:"MlbQciGroupId"`
	EmtcModeADrxParaGroupId string `xml:"EmtcModeADrxParaGroupId"`
	EmtcModeBDrxParaGroupId string `xml:"EmtcModeBDrxParaGroupId"`
	ServiceFlag string `xml:"ServiceFlag"`
	DlAmbrLimitFactor string `xml:"DlAmbrLimitFactor"`
	QciAmbrLimitFlag string `xml:"QciAmbrLimitFlag"`
	UlAmbrLimitFactor string `xml:"UlAmbrLimitFactor"`
}

