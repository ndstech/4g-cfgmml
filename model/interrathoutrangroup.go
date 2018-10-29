package model

import "encoding/xml"

type Interrathoutrangroup struct {
	XMLName xml.Name `xml:"INTERRATHOUTRANGROUP"`
	ATTRIBUTES InterrathoutrangroupAttributes `xml:"attributes"`
}

type InterrathoutrangroupAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	InterRatHoUtranGroupId string `xml:"InterRatHoUtranGroupId"`
	InterRatHoUtranB1ThdEcn0 string `xml:"InterRatHoUtranB1ThdEcn0"`
	InterRatHoUtranB1ThdRscp string `xml:"InterRatHoUtranB1ThdRscp"`
	InterRatHoUtranB1Hyst string `xml:"InterRatHoUtranB1Hyst"`
	InterRatHoUtranB1TimeToTrig string `xml:"InterRatHoUtranB1TimeToTrig"`
	LdSvBasedHoUtranB1ThdEcn0 string `xml:"LdSvBasedHoUtranB1ThdEcn0"`
	LdSvBasedHoUtranB1ThdRscp string `xml:"LdSvBasedHoUtranB1ThdRscp"`
	HSInterRatHoUtranB1TimeTrig string `xml:"HSInterRatHoUtranB1TimeTrig"`
}

