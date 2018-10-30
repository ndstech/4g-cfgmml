package model

import "encoding/xml"

type Pdschcfg struct {
	XMLName xml.Name `xml:"PDSCHCfg"`
	ATTRIBUTES PdschcfgAttributes `xml:"attributes"`
}

type PdschcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	ReferenceSignalPwr string `xml:"ReferenceSignalPwr"`
	Pb string `xml:"Pb"`
	ReferenceSignalPwrMargin string `xml:"ReferenceSignalPwrMargin"`
	TxPowerOffsetAnt0 string `xml:"TxPowerOffsetAnt0"`
	TxPowerOffsetAnt1 string `xml:"TxPowerOffsetAnt1"`
	TxPowerOffsetAnt2 string `xml:"TxPowerOffsetAnt2"`
	TxPowerOffsetAnt3 string `xml:"TxPowerOffsetAnt3"`
	TxChnPowerCfgSw string `xml:"TxChnPowerCfgSw"`
}

