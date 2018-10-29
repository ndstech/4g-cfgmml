package model

import "encoding/xml"

type Lwamgtcfg struct {
	XMLName xml.Name `xml:"LWAMGTCFG"`
	ATTRIBUTES LwamgtcfgAttributes `xml:"attributes"`
}

type LwamgtcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	LteWlanAggrW1ThdRssi string `xml:"LteWlanAggrW1ThdRssi"`
	LteWlanAggrW2Thd1Rssi string `xml:"LteWlanAggrW2Thd1Rssi"`
	LteWlanAggrW3ThdRssi string `xml:"LteWlanAggrW3ThdRssi"`
	LwaW1TimeToTrigger string `xml:"LwaW1TimeToTrigger"`
	LwaW2TimeToTrigger string `xml:"LwaW2TimeToTrigger"`
	LwaW3TimeToTrigger string `xml:"LwaW3TimeToTrigger"`
	WlanMeasHyst string `xml:"WlanMeasHyst"`
	CellLwaAlgoSwitch string `xml:"CellLwaAlgoSwitch"`
}

