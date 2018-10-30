package model

import "encoding/xml"

type Cellmimoparacfg struct {
	XMLName xml.Name `xml:"CellMimoParaCfg"`
	ATTRIBUTES CellmimoparacfgAttributes `xml:"attributes"`
}

type CellmimoparacfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	MimoAdaptiveSwitch string `xml:"MimoAdaptiveSwitch"`
	FixedMimoMode string `xml:"FixedMimoMode"`
	InitialMimoType string `xml:"InitialMimoType"`
}

