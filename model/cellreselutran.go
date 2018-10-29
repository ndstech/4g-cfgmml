package model

import "encoding/xml"

type Cellreselutran struct {
	XMLName xml.Name `xml:"CELLRESELUTRAN"`
	ATTRIBUTES CellreselutranAttributes `xml:"attributes"`
}

type CellreselutranAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	TReselUtran string `xml:"TReselUtran"`
	SpeedStateSfCfgInd string `xml:"SpeedStateSfCfgInd"`
	TReselUtranSfMedium string `xml:"TReselUtranSfMedium"`
	TReselUtranSfHigh string `xml:"TReselUtranSfHigh"`
}

