package model

import "encoding/xml"

type Cellucionpuschpara struct {
	XMLName xml.Name `xml:"CELLUCIONPUSCHPARA"`
	ATTRIBUTES CellucionpuschparaAttributes `xml:"attributes"`
}

type CellucionpuschparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	CellUciOnPuschParaValid string `xml:"CellUciOnPuschParaValid"`
	DeltaOffsetCqiIndex string `xml:"DeltaOffsetCqiIndex"`
	DeltaOffsetRiIndex string `xml:"DeltaOffsetRiIndex"`
	DeltaOffsetAckIndex string `xml:"DeltaOffsetAckIndex"`
	DeltaOffsetAckIndexForTtiB string `xml:"DeltaOffsetAckIndexForTtiB"`
}

