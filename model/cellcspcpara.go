package model

import "encoding/xml"

type Cellcspcpara struct {
	XMLName xml.Name `xml:"CellCspcPara"`
	ATTRIBUTES CellcspcparaAttributes `xml:"attributes"`
}

type CellcspcparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	ECspcPCAdjUeNumTh string `xml:"eCspcPCAdjUeNumTh"`
	CellCspcSwitch string `xml:"CellCspcSwitch"`
	CspcRapidRptSwitch string `xml:"CspcRapidRptSwitch"`
	CspcUeSrsCfgRptPeriod string `xml:"CspcUeSrsCfgRptPeriod"`
	UlRsrpRptPeriod string `xml:"UlRsrpRptPeriod"`
	CspcCqiFilterCoeff string `xml:"CspcCqiFilterCoeff"`
	UlRsrpPhyFilterCoeff string `xml:"UlRsrpPhyFilterCoeff"`
	UlRsrpRrmFilterCoeff string `xml:"UlRsrpRrmFilterCoeff"`
	ECspcA3Offset string `xml:"eCspcA3Offset"`
	CelleCspcSwitch string `xml:"CelleCspcSwitch"`
	ECspcPCAdjRange string `xml:"eCspcPCAdjRange"`
	IntraEnbCspcSw string `xml:"IntraEnbCspcSw"`
	ElasticCarrierSwitch string `xml:"ElasticCarrierSwitch"`
}

