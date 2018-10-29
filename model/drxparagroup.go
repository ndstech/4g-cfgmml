package model

import "encoding/xml"

type Drxparagroup struct {
	XMLName xml.Name `xml:"DRXPARAGROUP"`
	ATTRIBUTES DrxparagroupAttributes `xml:"attributes"`
}

type DrxparagroupAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	DrxParaGroupId string `xml:"DrxParaGroupId"`
	EnterDrxSwitch string `xml:"EnterDrxSwitch"`
	OnDurationTimer string `xml:"OnDurationTimer"`
	DrxInactivityTimer string `xml:"DrxInactivityTimer"`
	DrxReTxTimer string `xml:"DrxReTxTimer"`
	LongDrxCycle string `xml:"LongDrxCycle"`
	SupportShortDrx string `xml:"SupportShortDrx"`
	ShortDrxCycle string `xml:"ShortDrxCycle"`
	DrxShortCycleTimer string `xml:"DrxShortCycleTimer"`
	DrxUlReTxTimer string `xml:"DrxUlReTxTimer"`
	ExtendLongDrxCycleSwitch string `xml:"ExtendLongDrxCycleSwitch"`
	CatType string `xml:"CatType"`
}

