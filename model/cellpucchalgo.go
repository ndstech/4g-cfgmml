package model

import "encoding/xml"

type Cellpucchalgo struct {
	XMLName xml.Name `xml:"CELLPUCCHALGO"`
	ATTRIBUTES CellpucchalgoAttributes `xml:"attributes"`
}

type CellpucchalgoAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	SriLowLoadThd string `xml:"SriLowLoadThd"`
	SriReCfgInd string `xml:"SriReCfgInd"`
	SriAlgoSwitch string `xml:"SriAlgoSwitch"`
}

