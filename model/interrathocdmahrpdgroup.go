package model

import "encoding/xml"

type Interrathocdmahrpdgroup struct {
	XMLName xml.Name `xml:"INTERRATHOCDMAHRPDGROUP"`
	ATTRIBUTES InterrathocdmahrpdgroupAttributes `xml:"attributes"`
}

type InterrathocdmahrpdgroupAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	InterRatHoCdmaHrpdGroupId string `xml:"InterRatHoCdmaHrpdGroupId"`
	InterRatHoCdmaB1Hyst string `xml:"InterRatHoCdmaB1Hyst"`
	InterRatHoCdmaB1ThdPs string `xml:"InterRatHoCdmaB1ThdPs"`
	InterRatHoCdmaB1TimeToTrig string `xml:"InterRatHoCdmaB1TimeToTrig"`
	LdSvBasedHoCdmaB1ThdPs string `xml:"LdSvBasedHoCdmaB1ThdPs"`
	Cdma2000HrpdB2Thd1Rsrp string `xml:"Cdma2000HrpdB2Thd1Rsrp"`
	Cdma2000HrpdB2Thd1Rsrq string `xml:"Cdma2000HrpdB2Thd1Rsrq"`
	Cdma2000HrpdNonB2ThdRsrp string `xml:"Cdma2000HrpdNonB2ThdRsrp"`
	Cdma2000HrpdNonB2ThdRsrq string `xml:"Cdma2000HrpdNonB2ThdRsrq"`
}

