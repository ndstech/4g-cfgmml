package model

import "encoding/xml"

type Cellcecfg struct {
	XMLName xml.Name `xml:"CELLCECFG"`
	ATTRIBUTES CellcecfgAttributes `xml:"attributes"`
}

type CellcecfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	CoverageLevel string `xml:"CoverageLevel"`
	RachRsrpFstThd string `xml:"RachRsrpFstThd"`
	RachRsrpSndThd string `xml:"RachRsrpSndThd"`
	RachRsrpTrdThd string `xml:"RachRsrpTrdThd"`
}

