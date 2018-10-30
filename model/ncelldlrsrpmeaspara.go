package model

import "encoding/xml"

type Ncelldlrsrpmeaspara struct {
	XMLName xml.Name `xml:"NCellDlRsrpMeasPara"`
	ATTRIBUTES NcelldlrsrpmeasparaAttributes `xml:"attributes"`
}

type NcelldlrsrpmeasparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	DlRsrpAutoNCellMeasSwitch string `xml:"DlRsrpAutoNCellMeasSwitch"`
	NCellDlRsrpMeasA3Offset string `xml:"NCellDlRsrpMeasA3Offset"`
}

