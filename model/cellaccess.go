package model

import "encoding/xml"

type Cellaccess struct {
	XMLName xml.Name `xml:"CELLACCESS"`
	ATTRIBUTES CellaccessAttributes `xml:"attributes"`
}

type CellaccessAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	CellBarred string `xml:"CellBarred"`
	IntraFreqResel string `xml:"IntraFreqResel"`
	BroadcastMode string `xml:"BroadcastMode"`
	RoundPeriod string `xml:"RoundPeriod"`
	RoundActionTime string `xml:"RoundActionTime"`
	ReptSyncAvoidInd string `xml:"ReptSyncAvoidInd"`
	ReptSyncAvoidTime string `xml:"ReptSyncAvoidTime"`
}

