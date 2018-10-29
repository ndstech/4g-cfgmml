package model

import "encoding/xml"

type Cellulicicmcpara struct {
	XMLName xml.Name `xml:"CELLULICICMCPARA"`
	ATTRIBUTES CellulicicmcparaAttributes `xml:"attributes"`
}

type CellulicicmcparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	A3Offset string `xml:"A3Offset"`
}

