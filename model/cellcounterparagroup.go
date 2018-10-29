package model

import "encoding/xml"

type Cellcounterparagroup struct {
	XMLName xml.Name `xml:"CELLCOUNTERPARAGROUP"`
	ATTRIBUTES CellcounterparagroupAttributes `xml:"attributes"`
}

type CellcounterparagroupAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	CellUserNumPdfThd string `xml:"CellUserNumPdfThd"`
	UlIblerOptSwitch string `xml:"UlIblerOptSwitch"`
	CellCounterAlgoSwitch string `xml:"CellCounterAlgoSwitch"`
	EdgeUserA3Offset string `xml:"EdgeUserA3Offset"`
	DlThrpPdfThd0 string `xml:"DlThrpPdfThd0"`
	DlThrpPdfThd1 string `xml:"DlThrpPdfThd1"`
	DlThrpPdfThd2 string `xml:"DlThrpPdfThd2"`
	DlThrpPdfThd3 string `xml:"DlThrpPdfThd3"`
	DlThrpPdfThd4 string `xml:"DlThrpPdfThd4"`
	DlThrpPdfThd5 string `xml:"DlThrpPdfThd5"`
	DlThrpPdfThd6 string `xml:"DlThrpPdfThd6"`
	DlThrpPdfThd7 string `xml:"DlThrpPdfThd7"`
	DlThrpPdfThd8 string `xml:"DlThrpPdfThd8"`
	UlThrpPdfThd0 string `xml:"UlThrpPdfThd0"`
	UlThrpPdfThd1 string `xml:"UlThrpPdfThd1"`
	UlThrpPdfThd2 string `xml:"UlThrpPdfThd2"`
	UlThrpPdfThd3 string `xml:"UlThrpPdfThd3"`
	UlThrpPdfThd4 string `xml:"UlThrpPdfThd4"`
	UlThrpPdfThd5 string `xml:"UlThrpPdfThd5"`
	UlThrpPdfThd6 string `xml:"UlThrpPdfThd6"`
	UlThrpPdfThd7 string `xml:"UlThrpPdfThd7"`
	UlThrpPdfThd8 string `xml:"UlThrpPdfThd8"`
	UeAbnormalJudgeThd string `xml:"UeAbnormalJudgeThd"`
	ThrPdfPolicy string `xml:"ThrPdfPolicy"`
	EdgeUserServRSRPThd string `xml:"EdgeUserServRSRPThd"`
	EdgeUserSRSOffset string `xml:"EdgeUserSRSOffset"`
	ChrOutputModeSwitch string `xml:"ChrOutputModeSwitch"`
	DlUserXpUnsatiThd string `xml:"DlUserXpUnsatiThd"`
	DlTrafficVolumeThd string `xml:"DlTrafficVolumeThd"`
	DlPktDelayMeasPolicy string `xml:"DlPktDelayMeasPolicy"`
	QCI1ContinuousPktLossThld string `xml:"QCI1ContinuousPktLossThld"`
}

