package model

import "encoding/xml"

type Celldss struct {
	XMLName xml.Name `xml:"CELLDSS"`
	ATTRIBUTES CelldssAttributes `xml:"attributes"`
}

type CelldssAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	HighFreqShareRbNum string `xml:"HighFreqShareRbNum"`
	LowFreqShareRbNum string `xml:"LowFreqShareRbNum"`
	ReMuteSwitch string `xml:"ReMuteSwitch"`
	UlInterfRestrictionMode string `xml:"UlInterfRestrictionMode"`
	A3Offset string `xml:"A3Offset"`
	A6Offset string `xml:"A6Offset"`
	NearAreaSinrThd string `xml:"NearAreaSinrThd"`
	MiddleAreaSinrThd string `xml:"MiddleAreaSinrThd"`
	FarAreaSinrThd string `xml:"FarAreaSinrThd"`
	InterfNCellConfigMode string `xml:"InterfNCellConfigMode"`
	SpecShrPfmOptSwitch string `xml:"SpecShrPfmOptSwitch"`
	SinrThdWithoutGsmInterf string `xml:"SinrThdWithoutGsmInterf"`
	GsmInterfINThd string `xml:"GsmInterfINThd"`
	NearPtUserRsrpThd string `xml:"NearPtUserRsrpThd"`
}

