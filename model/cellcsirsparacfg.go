package model

import "encoding/xml"

type Cellcsirsparacfg struct {
	XMLName xml.Name `xml:"CellCsiRsParaCfg"`
	ATTRIBUTES CellcsirsparacfgAttributes `xml:"attributes"`
}

type CellcsirsparacfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	CsiRsSwitch string `xml:"CsiRsSwitch"`
	CsiRsPeriod string `xml:"CsiRsPeriod"`
	CsiRsConfigUserNumTh string `xml:"CsiRsConfigUserNumTh"`
	CsiRsUnconfigUserNumTh string `xml:"CsiRsUnconfigUserNumTh"`
	CsiRsConfigUserRatioTh string `xml:"CsiRsConfigUserRatioTh"`
	CsiRsUnconfigUserRatioTh string `xml:"CsiRsUnconfigUserRatioTh"`
	CsiRsSetJudgeHysTimer string `xml:"CsiRsSetJudgeHysTimer"`
	CsiRsSetJudgeTimer string `xml:"CsiRsSetJudgeTimer"`
	CsiRsPortNum string `xml:"CsiRsPortNum"`
}

