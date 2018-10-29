package model

import "encoding/xml"

type Eutranexternalcell struct {
	XMLName xml.Name `xml:"EUTRANEXTERNALCELL"`
	ATTRIBUTES EutranexternalcellAttributes `xml:"attributes"`
}

type EutranexternalcellAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	Mcc string `xml:"Mcc"`
	Mnc string `xml:"Mnc"`
	ENodeBId string `xml:"eNodeBId"`
	CellId string `xml:"CellId"`
	DlEarfcn string `xml:"DlEarfcn"`
	UlEarfcnCfgInd string `xml:"UlEarfcnCfgInd"`
	PhyCellId string `xml:"PhyCellId"`
	Tac string `xml:"Tac"`
	CellName string `xml:"CellName"`
	EutranExternalCellSlaveBand string `xml:"EutranExternalCellSlaveBand"`
	RoamingAreaHoInd string `xml:"RoamingAreaHoInd"`
	SpecifiedCellFlag string `xml:"SpecifiedCellFlag"`
	AnrFlag string `xml:"AnrFlag"`
	HighSpeedFlag string `xml:"HighSpeedFlag"`
	CtrlMode string `xml:"CtrlMode"`
	ObjId string `xml:"objId"`
	DlFreqOffset string `xml:"DlFreqOffset"`
	NbCellFlag string `xml:"NbCellFlag"`
	NclUpdateMode string `xml:"NclUpdateMode"`
	SupportEmtcFlag string `xml:"SupportEmtcFlag"`
}

