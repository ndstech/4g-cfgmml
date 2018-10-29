package model

import "encoding/xml"

type Cellmro struct {
	XMLName xml.Name `xml:"CELLMRO"`
	ATTRIBUTES CellmroAttributes `xml:"attributes"`
}

type CellmroAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	CioAdjLimitCfgInd string `xml:"CioAdjLimitCfgInd"`
	CioAdjUpperLimit string `xml:"CioAdjUpperLimit"`
	CioAdjLowerLimit string `xml:"CioAdjLowerLimit"`
	InterFreqA2RsrpUpLimit string `xml:"InterFreqA2RsrpUpLimit"`
	InterFreqA2RsrpLowLimit string `xml:"InterFreqA2RsrpLowLimit"`
	A3InterFreqA2RsrpUpLimit string `xml:"A3InterFreqA2RsrpUpLimit"`
	A3InterFreqA2RsrpLowLimit string `xml:"A3InterFreqA2RsrpLowLimit"`
	InterFreqMroAdjParaSel string `xml:"InterFreqMroAdjParaSel"`
}

