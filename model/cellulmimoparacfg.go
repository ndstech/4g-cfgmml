package model

import "encoding/xml"

type Cellulmimoparacfg struct {
	XMLName xml.Name `xml:"CellUlMimoParaCfg"`
	ATTRIBUTES CellulmimoparacfgAttributes `xml:"attributes"`
}

type CellulmimoparacfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	UlSuMimoRankPara string `xml:"UlSuMimoRankPara"`
}

