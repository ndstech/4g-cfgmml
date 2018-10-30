package model

import "encoding/xml"

type Interrathogerangroup struct {
	XMLName xml.Name `xml:"InterRatHoGeranGroup"`
	ATTRIBUTES InterrathogerangroupAttributes `xml:"attributes"`
}

type InterrathogerangroupAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	InterRatHoGeranGroupId string `xml:"InterRatHoGeranGroupId"`
	InterRatHoGeranB1Hyst string `xml:"InterRatHoGeranB1Hyst"`
	InterRatHoGeranB1Thd string `xml:"InterRatHoGeranB1Thd"`
	InterRatHoGeranB1TimeToTrig string `xml:"InterRatHoGeranB1TimeToTrig"`
	LdSvBasedHoGeranB1Thd string `xml:"LdSvBasedHoGeranB1Thd"`
}

