package model

import "encoding/xml"

type Cellulicalgo struct {
	XMLName xml.Name `xml:"CELLULICALGO"`
	ATTRIBUTES CellulicalgoAttributes `xml:"attributes"`
}

type CellulicalgoAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	UlIcA3Offset string `xml:"UlIcA3Offset"`
}

