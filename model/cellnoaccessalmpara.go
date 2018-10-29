package model

import "encoding/xml"

type Cellnoaccessalmpara struct {
	XMLName xml.Name `xml:"CELLNOACCESSALMPARA"`
	ATTRIBUTES CellnoaccessalmparaAttributes `xml:"attributes"`
}

type CellnoaccessalmparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	NoAccessStartDetectTime string `xml:"NoAccessStartDetectTime"`
	NoAccessStopDetectTime string `xml:"NoAccessStopDetectTime"`
	NoAccessDetectTimer string `xml:"NoAccessDetectTimer"`
}

