package model

import "encoding/xml"

type Cqiadaptivecfg struct {
	XMLName xml.Name `xml:"CQIADAPTIVECFG"`
	ATTRIBUTES CqiadaptivecfgAttributes `xml:"attributes"`
}

type CqiadaptivecfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CqiPeriodAdaptive string `xml:"CqiPeriodAdaptive"`
	SimulAckNackAndCqiSwitch string `xml:"SimulAckNackAndCqiSwitch"`
	PucchPeriodicCqiOptSwitch string `xml:"PucchPeriodicCqiOptSwitch"`
	HoAperiodicCqiCfgSwitch string `xml:"HoAperiodicCqiCfgSwitch"`
	SimulAckNackAndCqiFmt3Sw string `xml:"SimulAckNackAndCqiFmt3Sw"`
	ObjId string `xml:"objId"`
}

