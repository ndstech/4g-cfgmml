package model

import "encoding/xml"

type Geranncell struct {
	XMLName xml.Name `xml:"GERANNCELL"`
	ATTRIBUTES GeranncellAttributes `xml:"attributes"`
}

type GeranncellAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	Mcc string `xml:"Mcc"`
	Mnc string `xml:"Mnc"`
	Lac string `xml:"Lac"`
	GeranCellId string `xml:"GeranCellId"`
	NoRmvFlag string `xml:"NoRmvFlag"`
	NoHoFlag string `xml:"NoHoFlag"`
	BlindHoPriority string `xml:"BlindHoPriority"`
	AnrFlag string `xml:"AnrFlag"`
	LocalCellName string `xml:"LocalCellName"`
	NeighbourCellName string `xml:"NeighbourCellName"`
	OverlapInd string `xml:"OverlapInd"`
	NCellMeasPriority string `xml:"NCellMeasPriority"`
	CtrlMode string `xml:"CtrlMode"`
}

