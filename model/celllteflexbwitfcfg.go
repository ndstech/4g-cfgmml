package model

import "encoding/xml"

type Celllteflexbwitfcfg struct {
	XMLName xml.Name `xml:"CellLteFlexBwItfCfg"`
	ATTRIBUTES CelllteflexbwitfcfgAttributes `xml:"attributes"`
}

type CelllteflexbwitfcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	HighItfGsmArfcn string `xml:"HighItfGsmArfcn"`
	GsmCarrierFreq string `xml:"GsmCarrierFreq"`
}

