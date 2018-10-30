package model

import "encoding/xml"

type Interrathocdma1xrttgroup struct {
	XMLName xml.Name `xml:"InterRatHoCdma1xRttGroup"`
	ATTRIBUTES Interrathocdma1xrttgroupAttributes `xml:"attributes"`
}

type Interrathocdma1xrttgroupAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	InterRatHoCdma1xRttGroupId string `xml:"InterRatHoCdma1xRttGroupId"`
	InterRatHoCdmaB1Hyst string `xml:"InterRatHoCdmaB1Hyst"`
	InterRatHoCdmaB1ThdPs string `xml:"InterRatHoCdmaB1ThdPs"`
	InterRatHoCdmaB1TimeToTrig string `xml:"InterRatHoCdmaB1TimeToTrig"`
	LdSvBasedHoCdmaB1ThdPs string `xml:"LdSvBasedHoCdmaB1ThdPs"`
}

