package model

import "encoding/xml"

type Cellceschcfg struct {
	XMLName xml.Name `xml:"CellCeSchCfg"`
	ATTRIBUTES CellceschcfgAttributes `xml:"attributes"`
}

type CellceschcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	CoverageLevel string `xml:"CoverageLevel"`
	SIB1RepNum string `xml:"SIB1RepNum"`
	PdschMaxNumRepModeA string `xml:"PdschMaxNumRepModeA"`
	PdschMaxNumRepModeB string `xml:"PdschMaxNumRepModeB"`
	PuschMaxNumRepModeA string `xml:"PuschMaxNumRepModeA"`
	PuschMaxNumRepModeB string `xml:"PuschMaxNumRepModeB"`
	MpdcchMaxNumRepPaging string `xml:"MpdcchMaxNumRepPaging"`
	MpdcchMaxNumRepModeA string `xml:"MpdcchMaxNumRepModeA"`
	MpdcchMaxNumRepModeB string `xml:"MpdcchMaxNumRepModeB"`
	SiTransEcr string `xml:"SiTransEcr"`
	PagingGroupNum string `xml:"PagingGroupNum"`
}

