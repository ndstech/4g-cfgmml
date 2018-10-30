package model

import "encoding/xml"

type Simuload struct {
	XMLName xml.Name `xml:"SimuLoad"`
	ATTRIBUTES SimuloadAttributes `xml:"attributes"`
}

type SimuloadAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	SimuLoadCfgIndex string `xml:"SimuLoadCfgIndex"`
	SimuLoadRbThd string `xml:"SimuLoadRbThd"`
	SimuLoadPwrThd string `xml:"SimuLoadPwrThd"`
	SimuLoadReportPeriod string `xml:"SimuLoadReportPeriod"`
	SimuLoadStatPeriod string `xml:"SimuLoadStatPeriod"`
	FreqSelSwitch string `xml:"FreqSelSwitch"`
	ObjId string `xml:"objId"`
}

