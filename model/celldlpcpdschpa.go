package model

import "encoding/xml"

type Celldlpcpdschpa struct {
	XMLName xml.Name `xml:"CellDlpcPdschPa"`
	ATTRIBUTES CelldlpcpdschpaAttributes `xml:"attributes"`
}

type CelldlpcpdschpaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	PdschPaAdjSwitch string `xml:"PdschPaAdjSwitch"`
	PaPcOff string `xml:"PaPcOff"`
	NomPdschRsEpreOffset string `xml:"NomPdschRsEpreOffset"`
}

