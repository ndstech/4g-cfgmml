package model

import "encoding/xml"

type Cellsrsadaptivecfg struct {
	XMLName xml.Name `xml:"CELLSRSADAPTIVECFG"`
	ATTRIBUTES CellsrsadaptivecfgAttributes `xml:"attributes"`
}

type CellsrsadaptivecfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	SrsPeriodAdaptive string `xml:"SrsPeriodAdaptive"`
}

