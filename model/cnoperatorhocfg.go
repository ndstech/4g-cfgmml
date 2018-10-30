package model

import "encoding/xml"

type Cnoperatorhocfg struct {
	XMLName xml.Name `xml:"CnOperatorHoCfg"`
	ATTRIBUTES CnoperatorhocfgAttributes `xml:"attributes"`
}

type CnoperatorhocfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CnOperatorId string `xml:"CnOperatorId"`
	FirstRatPri string `xml:"FirstRatPri"`
	SecondRatPri string `xml:"SecondRatPri"`
	TddIfHoA2ThdRsrpOffset string `xml:"TddIfHoA2ThdRsrpOffset"`
	FddIfHoA2ThdRsrpOffset string `xml:"FddIfHoA2ThdRsrpOffset"`
	UtranA2ThdRsrpOffset string `xml:"UtranA2ThdRsrpOffset"`
	GeranA2ThdRsrpOffset string `xml:"GeranA2ThdRsrpOffset"`
	SrvccWithPsBasedCellCapSw string `xml:"SrvccWithPsBasedCellCapSw"`
	PsInterRatSecondPri string `xml:"PsInterRatSecondPri"`
	PsInterRatHighestPri string `xml:"PsInterRatHighestPri"`
	PsInterRatLowestPri string `xml:"PsInterRatLowestPri"`
}

