package model

import "encoding/xml"

type Blindncellopt struct {
	XMLName xml.Name `xml:"BLINDNCELLOPT"`
	ATTRIBUTES BlindncelloptAttributes `xml:"attributes"`
}

type BlindncelloptAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	StatisticPeriod string `xml:"StatisticPeriod"`
	SampleNumThd string `xml:"SampleNumThd"`
	HoSuccRateThd string `xml:"HoSuccRateThd"`
	CsfbHoAttempRatioThd string `xml:"CsfbHoAttempRatioThd"`
	BlindHoSuccRateThd string `xml:"BlindHoSuccRateThd"`
	OptMode string `xml:"OptMode"`
	ObjId string `xml:"objId"`
}

