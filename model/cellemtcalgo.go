package model

import "encoding/xml"

type Cellemtcalgo struct {
	XMLName xml.Name `xml:"CellEmtcAlgo"`
	ATTRIBUTES CellemtcalgoAttributes `xml:"attributes"`
}

type CellemtcalgoAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	EmtcUlRbTargetRatio string `xml:"EmtcUlRbTargetRatio"`
	EmtcAperCqiTrigPrd string `xml:"EmtcAperCqiTrigPrd"`
	EmtcDlRbTargetRatio string `xml:"EmtcDlRbTargetRatio"`
	EmtcAlgoSwitch string `xml:"EmtcAlgoSwitch"`
	DlLteRsvNbCount string `xml:"DlLteRsvNbCount"`
	UlLteRsvNbCount string `xml:"UlLteRsvNbCount"`
	EmtcVolteSupportSwitch string `xml:"EmtcVolteSupportSwitch"`
}

