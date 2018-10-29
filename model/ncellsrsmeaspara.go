package model

import "encoding/xml"

type Ncellsrsmeaspara struct {
	XMLName xml.Name `xml:"NCELLSRSMEASPARA"`
	ATTRIBUTES NcellsrsmeasparaAttributes `xml:"attributes"`
}

type NcellsrsmeasparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	SrsAutoNCellMeasSwitch string `xml:"SrsAutoNCellMeasSwitch"`
	NCellSrsMeasA3Offset string `xml:"NCellSrsMeasA3Offset"`
	NCellMeasSwitch string `xml:"NCellMeasSwitch"`
}

