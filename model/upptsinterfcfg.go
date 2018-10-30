package model

import "encoding/xml"

type Upptsinterfcfg struct {
	XMLName xml.Name `xml:"UpPTSInterfCfg"`
	ATTRIBUTES UpptsinterfcfgAttributes `xml:"attributes"`
}

type UpptsinterfcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	TestPeriod string `xml:"TestPeriod"`
	TestThreshold string `xml:"TestThreshold"`
	TestHysteresis string `xml:"TestHysteresis"`
}

