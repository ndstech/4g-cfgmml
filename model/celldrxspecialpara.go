package model

import "encoding/xml"

type Celldrxspecialpara struct {
	XMLName xml.Name `xml:"CELLDRXSPECIALPARA"`
	ATTRIBUTES CelldrxspecialparaAttributes `xml:"attributes"`
}

type CelldrxspecialparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	CellDrxSpecialParaValid string `xml:"CellDrxSpecialParaValid"`
	LongDrxCycleSpecial string `xml:"LongDrxCycleSpecial"`
	OnDurationTimerSpecial string `xml:"OnDurationTimerSpecial"`
	DrxInactivityTimerSpecial string `xml:"DrxInactivityTimerSpecial"`
	ShortDrxCycleSwitchSpecial string `xml:"ShortDrxCycleSwitchSpecial"`
	LongDrxCycleForIntraRatAnr string `xml:"LongDrxCycleForIntraRatAnr"`
	LongDrxCycleForInterRatAnr string `xml:"LongDrxCycleForInterRatAnr"`
	FddAnrDrxInactivityTimer string `xml:"FddAnrDrxInactivityTimer"`
	TddAnrDrxInactivityTimer string `xml:"TddAnrDrxInactivityTimer"`
	BfSpecificDrxParaSwitch string `xml:"BfSpecificDrxParaSwitch"`
}

