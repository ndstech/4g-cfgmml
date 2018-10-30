package model

import "encoding/xml"

type Cellsel struct {
	XMLName xml.Name `xml:"CellSel"`
	ATTRIBUTES CellselAttributes `xml:"attributes"`
}

type CellselAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	QRxLevMin string `xml:"QRxLevMin"`
	QRxLevMinOffset string `xml:"QRxLevMinOffset"`
	QQualMin string `xml:"QQualMin"`
	QQualMinOffsetCfgInd string `xml:"QQualMinOffsetCfgInd"`
	QRxLevMinCE string `xml:"QRxLevMinCE"`
	QQualMinCE string `xml:"QQualMinCE"`
}

