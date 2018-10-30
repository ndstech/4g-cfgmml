package model

import "encoding/xml"

type Cellwttxparacfg struct {
	XMLName xml.Name `xml:"CellWttxParaCfg"`
	ATTRIBUTES CellwttxparacfgAttributes `xml:"attributes"`
}

type CellwttxparacfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	MbbUserDlPrbUpLimit string `xml:"MbbUserDlPrbUpLimit"`
	MbbUserUlPrbUpLimit string `xml:"MbbUserUlPrbUpLimit"`
	PrbUpLimitCtrlMode string `xml:"PrbUpLimitCtrlMode"`
	WbbOrMbbUserDefMode string `xml:"WbbOrMbbUserDefMode"`
	WbbUserDlPrbUpLimit string `xml:"WbbUserDlPrbUpLimit"`
	WbbUserUlPrbUpLimit string `xml:"WbbUserUlPrbUpLimit"`
}

