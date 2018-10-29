package model

import "encoding/xml"

type Cellreselgeran struct {
	XMLName xml.Name `xml:"CELLRESELGERAN"`
	ATTRIBUTES CellreselgeranAttributes `xml:"attributes"`
}

type CellreselgeranAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	TReselGeran string `xml:"TReselGeran"`
	SpeedStateSfCfgInd string `xml:"SpeedStateSfCfgInd"`
}

