package model

import "encoding/xml"

type Cellcqiadaptivecfg struct {
	XMLName xml.Name `xml:"CELLCQIADAPTIVECFG"`
	ATTRIBUTES CellcqiadaptivecfgAttributes `xml:"attributes"`
}

type CellcqiadaptivecfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	CqiPeriodAdaptive string `xml:"CqiPeriodAdaptive"`
	SimulAckNackAndCqiSwitch string `xml:"SimulAckNackAndCqiSwitch"`
	HoAperiodicCqiCfgSwitch string `xml:"HoAperiodicCqiCfgSwitch"`
	MinCqiPeriod string `xml:"MinCqiPeriod"`
	SccCqiRptEnhancedSwitch string `xml:"SccCqiRptEnhancedSwitch"`
}

