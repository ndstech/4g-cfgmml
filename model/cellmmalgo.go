package model

import "encoding/xml"

type Cellmmalgo struct {
	XMLName xml.Name `xml:"CELLMMALGO"`
	ATTRIBUTES CellmmalgoAttributes `xml:"attributes"`
}

type CellmmalgoAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	BeamGainLimitSwitch string `xml:"BeamGainLimitSwitch"`
	MMAlgoOptSwitch string `xml:"MMAlgoOptSwitch"`
}

