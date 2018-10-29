package model

import "encoding/xml"

type Cellsrlte struct {
	XMLName xml.Name `xml:"CELLSRLTE"`
	ATTRIBUTES CellsrlteAttributes `xml:"attributes"`
}

type CellsrlteAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	SrlteSwitch string `xml:"SrlteSwitch"`
	SrlteDtxThrd string `xml:"SrlteDtxThrd"`
	SrlteSuspendTime string `xml:"SrlteSuspendTime"`
	SrlteDiscardAlgoSwitch string `xml:"SrlteDiscardAlgoSwitch"`
}

