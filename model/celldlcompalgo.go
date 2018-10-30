package model

import "encoding/xml"

type Celldlcompalgo struct {
	XMLName xml.Name `xml:"CellDlCompAlgo"`
	ATTRIBUTES CelldlcompalgoAttributes `xml:"attributes"`
}

type CelldlcompalgoAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	DlCompA3Offset string `xml:"DlCompA3Offset"`
	DpsCoCellUserRatioThd string `xml:"DpsCoCellUserRatioThd"`
	DpsLoadDiffThd string `xml:"DpsLoadDiffThd"`
	DpsServingCellDlPrbThd string `xml:"DpsServingCellDlPrbThd"`
	HetnetDpsCoCellRsrpThd string `xml:"HetnetDpsCoCellRsrpThd"`
	HomnetDpsCoCellRsrpThd string `xml:"HomnetDpsCoCellRsrpThd"`
	JtCoCellDlPrbThd string `xml:"JtCoCellDlPrbThd"`
	JtCoCellRsrpThd string `xml:"JtCoCellRsrpThd"`
	JtServingCellDlPrbThd string `xml:"JtServingCellDlPrbThd"`
}

