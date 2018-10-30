package model

import "encoding/xml"

type Cellulpcdedic struct {
	XMLName xml.Name `xml:"CellUlpcDedic"`
	ATTRIBUTES CellulpcdedicAttributes `xml:"attributes"`
}

type CellulpcdedicAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	DeltaMcsEnabled string `xml:"DeltaMcsEnabled"`
	PSrsOffsetDeltaMcsDisable string `xml:"PSrsOffsetDeltaMcsDisable"`
	PSrsOffsetDeltaMcsEnable string `xml:"PSrsOffsetDeltaMcsEnable"`
	FilterRsrp string `xml:"FilterRsrp"`
	PathlossReferenceLink string `xml:"PathlossReferenceLink"`
}

