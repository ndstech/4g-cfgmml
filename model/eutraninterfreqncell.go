package model

import "encoding/xml"

type Eutraninterfreqncell struct {
	XMLName xml.Name `xml:"EutranInterFreqNCell"`
	ATTRIBUTES EutraninterfreqncellAttributes `xml:"attributes"`
}

type EutraninterfreqncellAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	Mcc string `xml:"Mcc"`
	Mnc string `xml:"Mnc"`
	ENodeBId string `xml:"eNodeBId"`
	CellId string `xml:"CellId"`
	CellIndividualOffset string `xml:"CellIndividualOffset"`
	CellQoffset string `xml:"CellQoffset"`
	NoHoFlag string `xml:"NoHoFlag"`
	NoRmvFlag string `xml:"NoRmvFlag"`
	BlindHoPriority string `xml:"BlindHoPriority"`
	AnrFlag string `xml:"AnrFlag"`
	LocalCellName string `xml:"LocalCellName"`
	NeighbourCellName string `xml:"NeighbourCellName"`
	CellMeasPriority string `xml:"CellMeasPriority"`
	OverlapInd string `xml:"OverlapInd"`
	OverlapRange string `xml:"OverlapRange"`
	NCellClassLabel string `xml:"NCellClassLabel"`
	CtrlMode string `xml:"CtrlMode"`
	AggregationProperty string `xml:"AggregationProperty"`
	OverlapIndicatorExtension string `xml:"OverlapIndicatorExtension"`
}

