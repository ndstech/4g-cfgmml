package model

import "encoding/xml"

type Csfallbackblindhocfg struct {
	XMLName xml.Name `xml:"CSFallBackBlindHoCfg"`
	ATTRIBUTES CsfallbackblindhocfgAttributes `xml:"attributes"`
}

type CsfallbackblindhocfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CnOperatorId string `xml:"CnOperatorId"`
	InterRatHighestPri string `xml:"InterRatHighestPri"`
	InterRatSecondPri string `xml:"InterRatSecondPri"`
	InterRatLowestPri string `xml:"InterRatLowestPri"`
	UtranLcsCap string `xml:"UtranLcsCap"`
	GeranLcsCap string `xml:"GeranLcsCap"`
	CdmaLcsCap string `xml:"CdmaLcsCap"`
	IdleCsfbHighestPri string `xml:"IdleCsfbHighestPri"`
	IdleCsfbSecondPri string `xml:"IdleCsfbSecondPri"`
	IdleCsfbLowestPri string `xml:"IdleCsfbLowestPri"`
	UtranCsfbBlindRedirRrSw string `xml:"UtranCsfbBlindRedirRrSw"`
}

