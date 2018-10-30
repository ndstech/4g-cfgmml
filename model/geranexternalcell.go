package model

import "encoding/xml"

type Geranexternalcell struct {
	XMLName xml.Name `xml:"GeranExternalCell"`
	ATTRIBUTES GeranexternalcellAttributes `xml:"attributes"`
}

type GeranexternalcellAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	Mcc string `xml:"Mcc"`
	Mnc string `xml:"Mnc"`
	GeranCellId string `xml:"GeranCellId"`
	Lac string `xml:"Lac"`
	RacCfgInd string `xml:"RacCfgInd"`
	Rac string `xml:"Rac"`
	BandIndicator string `xml:"BandIndicator"`
	GeranArfcn string `xml:"GeranArfcn"`
	NetworkColourCode string `xml:"NetworkColourCode"`
	BaseStationColourCode string `xml:"BaseStationColourCode"`
	DtmInd string `xml:"DtmInd"`
	CellName string `xml:"CellName"`
	CsPsHOInd string `xml:"CsPsHOInd"`
	UltraFlashCsfbInd string `xml:"UltraFlashCsfbInd"`
	RoamingAreaHoInd string `xml:"RoamingAreaHoInd"`
	AnrFlag string `xml:"AnrFlag"`
	CtrlMode string `xml:"CtrlMode"`
	ObjId string `xml:"objId"`
}

