package model

import "encoding/xml"

type Utranranshare struct {
	XMLName xml.Name `xml:"UtranRanShare"`
	ATTRIBUTES UtranranshareAttributes `xml:"attributes"`
}

type UtranranshareAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	UtranDlArfcn string `xml:"UtranDlArfcn"`
	Mcc string `xml:"Mcc"`
	Mnc string `xml:"Mnc"`
	CellReselPriorityCfgInd string `xml:"CellReselPriorityCfgInd"`
}

