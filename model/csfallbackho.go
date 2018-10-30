package model

import "encoding/xml"

type Csfallbackho struct {
	XMLName xml.Name `xml:"CSFallBackHo"`
	ATTRIBUTES CsfallbackhoAttributes `xml:"attributes"`
}

type CsfallbackhoAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	CsfbHoUtranTimeToTrig string `xml:"CsfbHoUtranTimeToTrig"`
	CsfbHoGeranTimeToTrig string `xml:"CsfbHoGeranTimeToTrig"`
	CsfbHoCdmaTimeToTrig string `xml:"CsfbHoCdmaTimeToTrig"`
	CsfbHoUtranB1ThdRscp string `xml:"CsfbHoUtranB1ThdRscp"`
	CsfbHoUtranB1ThdEcn0 string `xml:"CsfbHoUtranB1ThdEcn0"`
	CsfbHoGeranB1Thd string `xml:"CsfbHoGeranB1Thd"`
	CsfbHoCdmaB1ThdPs string `xml:"CsfbHoCdmaB1ThdPs"`
	BlindHoA1ThdRsrp string `xml:"BlindHoA1ThdRsrp"`
	CsfbProtectionTimer string `xml:"CsfbProtectionTimer"`
	CsfbProtectTimer string `xml:"CsfbProtectTimer"`
}

