package model

import "encoding/xml"

type Eutranintrafreqncell struct {
	XMLName xml.Name `xml:"EutranIntraFreqNCell"`
	ATTRIBUTES EutranintrafreqncellAttributes `xml:"attributes"`
}

type EutranintrafreqncellAttributes struct {
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
	AnrFlag string `xml:"AnrFlag"`
	LocalCellName string `xml:"LocalCellName"`
	NeighbourCellName string `xml:"NeighbourCellName"`
	CellMeasPriority string `xml:"CellMeasPriority"`
	CellRangeExpansion string `xml:"CellRangeExpansion"`
	NCellClassLabel string `xml:"NCellClassLabel"`
	CtrlMode string `xml:"CtrlMode"`
	AttachCellSwitch string `xml:"AttachCellSwitch"`
	HighSpeedCellIndOffset string `xml:"HighSpeedCellIndOffset"`
	VectorCellFlag string `xml:"VectorCellFlag"`
}

