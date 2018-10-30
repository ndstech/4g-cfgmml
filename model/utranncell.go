package model

import "encoding/xml"

type Utranncell struct {
	XMLName xml.Name `xml:"UtranNCell"`
	ATTRIBUTES UtranncellAttributes `xml:"attributes"`
}

type UtranncellAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	Mcc string `xml:"Mcc"`
	Mnc string `xml:"Mnc"`
	RncId string `xml:"RncId"`
	CellId string `xml:"CellId"`
	NoHoFlag string `xml:"NoHoFlag"`
	NoRmvFlag string `xml:"NoRmvFlag"`
	AnrFlag string `xml:"AnrFlag"`
	BlindHoPriority string `xml:"BlindHoPriority"`
	CellMeasPriority string `xml:"CellMeasPriority"`
	OverlapInd string `xml:"OverlapInd"`
	NCellMeasPriority string `xml:"NCellMeasPriority"`
	LocalCellName string `xml:"LocalCellName"`
	NeighbourCellName string `xml:"NeighbourCellName"`
	CtrlMode string `xml:"CtrlMode"`
}

