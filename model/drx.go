package model

import "encoding/xml"

type Drx struct {
	XMLName xml.Name `xml:"DRX"`
	ATTRIBUTES DrxAttributes `xml:"attributes"`
}

type DrxAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	DrxAlgSwitch string `xml:"DrxAlgSwitch"`
	ShortDrxSwitch string `xml:"ShortDrxSwitch"`
	LongDrxCycleSpecial string `xml:"LongDrxCycleSpecial"`
	OnDurationTimerSpecial string `xml:"OnDurationTimerSpecial"`
	DrxInactivityTimerSpecial string `xml:"DrxInactivityTimerSpecial"`
	SupportShortDrxSpecial string `xml:"SupportShortDrxSpecial"`
	LongDrxCycleForAnr string `xml:"LongDrxCycleForAnr"`
	LongDRXCycleforIRatAnr string `xml:"LongDRXCycleforIRatAnr"`
	DrxInactivityTimerForAnr string `xml:"DrxInactivityTimerForAnr"`
	TddAnrDrxInactivityTimer string `xml:"TddAnrDrxInactivityTimer"`
	ObjId string `xml:"objId"`
	DrxFlexCtrlSwitch string `xml:"DrxFlexCtrlSwitch"`
}

