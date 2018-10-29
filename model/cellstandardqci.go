package model

import "encoding/xml"

type Cellstandardqci struct {
	XMLName xml.Name `xml:"CELLSTANDARDQCI"`
	ATTRIBUTES CellstandardqciAttributes `xml:"attributes"`
}

type CellstandardqciAttributes struct {
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
	RlfTimerConstGroupId string `xml:"RlfTimerConstGroupId"`
	TrafficRelDelay string `xml:"TrafficRelDelay"`
	QciPriorityForHo string `xml:"QciPriorityForHo"`
	MlbQciGroupId string `xml:"MlbQciGroupId"`
	PreallocationParaGroupId string `xml:"PreallocationParaGroupId"`
	QciPriorityForDrx string `xml:"QciPriorityForDrx"`
	QciAlgoSwitch string `xml:"QciAlgoSwitch"`
	QciEutranFreqRelationId string `xml:"QciEutranFreqRelationId"`
	QciUtranFreqRelationId string `xml:"QciUtranFreqRelationId"`
	QciGeranFreqRelationId string `xml:"QciGeranFreqRelationId"`
	QciAmbrLimitFlag string `xml:"QciAmbrLimitFlag"`
}

