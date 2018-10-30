package model

import "encoding/xml"

type Intrafreqhogroup struct {
	XMLName xml.Name `xml:"IntraFreqHoGroup"`
	ATTRIBUTES IntrafreqhogroupAttributes `xml:"attributes"`
}

type IntrafreqhogroupAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	IntraFreqHoGroupId string `xml:"IntraFreqHoGroupId"`
	IntraFreqHoA3Hyst string `xml:"IntraFreqHoA3Hyst"`
	IntraFreqHoA3Offset string `xml:"IntraFreqHoA3Offset"`
	IntraFreqHoA3TimeToTrig string `xml:"IntraFreqHoA3TimeToTrig"`
	HighSpeedA3TimeToTrig string `xml:"HighSpeedA3TimeToTrig"`
	IntraFreqHoA2ThldRsrpCE string `xml:"IntraFreqHoA2ThldRsrpCE"`
}

