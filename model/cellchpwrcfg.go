package model

import "encoding/xml"

type Cellchpwrcfg struct {
	XMLName xml.Name `xml:"CellChPwrCfg"`
	ATTRIBUTES CellchpwrcfgAttributes `xml:"attributes"`
}

type CellchpwrcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	PcfichPwr string `xml:"PcfichPwr"`
	PbchPwr string `xml:"PbchPwr"`
	SchPwr string `xml:"SchPwr"`
	DbchPwr string `xml:"DbchPwr"`
	PchPwr string `xml:"PchPwr"`
	RaRspPwr string `xml:"RaRspPwr"`
	PrsPwr string `xml:"PrsPwr"`
	AntOutputPwr string `xml:"AntOutputPwr"`
	PmchPwrOffset string `xml:"PmchPwrOffset"`
	HoRarPwrEnhancedSwitch string `xml:"HoRarPwrEnhancedSwitch"`
}

