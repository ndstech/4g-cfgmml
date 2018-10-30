package model

import "encoding/xml"

type Cellmlbuesel struct {
	XMLName xml.Name `xml:"CellMlbUeSel"`
	ATTRIBUTES CellmlbueselAttributes `xml:"attributes"`
}

type CellmlbueselAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	UeSelectPrbPrio string `xml:"UeSelectPrbPrio"`
	UeSelectQciPrio string `xml:"UeSelectQciPrio"`
	UeSelectArpPrio string `xml:"UeSelectArpPrio"`
	UeSelectDlMcsPrio string `xml:"UeSelectDlMcsPrio"`
	InterFreqMlbUeArpThd string `xml:"InterFreqMlbUeArpThd"`
	InterFreqMlbUeDlMcsThd string `xml:"InterFreqMlbUeDlMcsThd"`
}

