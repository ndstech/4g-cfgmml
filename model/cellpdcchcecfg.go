package model

import "encoding/xml"

type Cellpdcchcecfg struct {
	XMLName xml.Name `xml:"CELLPDCCHCECFG"`
	ATTRIBUTES CellpdcchcecfgAttributes `xml:"attributes"`
}

type CellpdcchcecfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	CoverageLevel string `xml:"CoverageLevel"`
	PdcchMaxRepetitionCnt string `xml:"PdcchMaxRepetitionCnt"`
	PdcchPeriodFactor string `xml:"PdcchPeriodFactor"`
	PdcchTransRptCntFactor string `xml:"PdcchTransRptCntFactor"`
}

