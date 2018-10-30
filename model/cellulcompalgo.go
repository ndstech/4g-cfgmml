package model

import "encoding/xml"

type Cellulcompalgo struct {
	XMLName xml.Name `xml:"CellUlCompAlgo"`
	ATTRIBUTES CellulcompalgoAttributes `xml:"attributes"`
}

type CellulcompalgoAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	UlCompA3Offset string `xml:"UlCompA3Offset"`
	SfnUlCompThd string `xml:"SfnUlCompThd"`
	UlCompA3OffsetForRelaxedBH string `xml:"UlCompA3OffsetForRelaxedBH"`
	UlCompUlRsrpOffset string `xml:"UlCompUlRsrpOffset"`
}

