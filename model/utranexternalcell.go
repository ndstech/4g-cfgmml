package model

import "encoding/xml"

type Utranexternalcell struct {
	XMLName xml.Name `xml:"UTRANEXTERNALCELL"`
	ATTRIBUTES UtranexternalcellAttributes `xml:"attributes"`
}

type UtranexternalcellAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	Mcc string `xml:"Mcc"`
	Mnc string `xml:"Mnc"`
	RncId string `xml:"RncId"`
	CellId string `xml:"CellId"`
	UtranDlArfcn string `xml:"UtranDlArfcn"`
	UtranUlArfcnCfgInd string `xml:"UtranUlArfcnCfgInd"`
	UtranFddTddType string `xml:"UtranFddTddType"`
	RacCfgInd string `xml:"RacCfgInd"`
	Rac string `xml:"Rac"`
	PScrambCode string `xml:"PScrambCode"`
	Lac string `xml:"Lac"`
	CellName string `xml:"CellName"`
	CsPsHOInd string `xml:"CsPsHOInd"`
	UtranExternalCellSlaveBand string `xml:"UtranExternalCellSlaveBand"`
	RoamingAreaHoInd string `xml:"RoamingAreaHoInd"`
	AnrFlag string `xml:"AnrFlag"`
	CtrlMode string `xml:"CtrlMode"`
	ObjId string `xml:"objId"`
	UtranUlArfcn string `xml:"UtranUlArfcn"`
}

