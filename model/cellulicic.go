package model

import "encoding/xml"

type Cellulicic struct {
	XMLName xml.Name `xml:"CELLULICIC"`
	ATTRIBUTES CellulicicAttributes `xml:"attributes"`
}

type CellulicicAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	BandMode string `xml:"BandMode"`
}

