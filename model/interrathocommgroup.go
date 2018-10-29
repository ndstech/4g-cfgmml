package model

import "encoding/xml"

type Interrathocommgroup struct {
	XMLName xml.Name `xml:"INTERRATHOCOMMGROUP"`
	ATTRIBUTES InterrathocommgroupAttributes `xml:"attributes"`
}

type InterrathocommgroupAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	InterRatHoCommGroupId string `xml:"InterRatHoCommGroupId"`
	InterRatHoA1A2Hyst string `xml:"InterRatHoA1A2Hyst"`
	InterRatHoA1A2TimeToTrig string `xml:"InterRatHoA1A2TimeToTrig"`
	InterRatHoA1ThdRsrp string `xml:"InterRatHoA1ThdRsrp"`
	InterRatHoA1ThdRsrq string `xml:"InterRatHoA1ThdRsrq"`
	InterRatHoA2ThdRsrp string `xml:"InterRatHoA2ThdRsrp"`
	InterRatHoA2ThdRsrq string `xml:"InterRatHoA2ThdRsrq"`
	GeranB2Thd1Rsrp string `xml:"GeranB2Thd1Rsrp"`
	GeranB2Thd1Rsrq string `xml:"GeranB2Thd1Rsrq"`
	UtranB2Thd1Rsrp string `xml:"UtranB2Thd1Rsrp"`
	UtranB2Thd1Rsrq string `xml:"UtranB2Thd1Rsrq"`
	DelIfMeasA2ThdRsrpOffset string `xml:"DelIfMeasA2ThdRsrpOffset"`
	DelIfMeasA2ThdRsrqOffset string `xml:"DelIfMeasA2ThdRsrqOffset"`
}

