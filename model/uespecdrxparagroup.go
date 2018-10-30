package model

import "encoding/xml"

type Uespecdrxparagroup struct {
	XMLName xml.Name `xml:"UeSpecDrxParaGroup"`
	ATTRIBUTES UespecdrxparagroupAttributes `xml:"attributes"`
}

type UespecdrxparagroupAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	DrxParaGroupIndex string `xml:"DrxParaGroupIndex"`
	DrxParaGroupStat string `xml:"DrxParaGroupStat"`
	OnDurationTimer string `xml:"OnDurationTimer"`
	DrxInactivityTimer string `xml:"DrxInactivityTimer"`
	DrxReTxTimer string `xml:"DrxReTxTimer"`
	LongDrxCycle string `xml:"LongDrxCycle"`
	SupportShortDrx string `xml:"SupportShortDrx"`
	ShortDrxCycle string `xml:"ShortDrxCycle"`
	DrxShortCycleTimer string `xml:"DrxShortCycleTimer"`
}

