package model

import "encoding/xml"

type Cellrachalgo struct {
	XMLName xml.Name `xml:"CELLRACHALGO"`
	ATTRIBUTES CellrachalgoAttributes `xml:"attributes"`
}

type CellrachalgoAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	PrachFalseAlarmDetRadThd string `xml:"PrachFalseAlarmDetRadThd"`
	RachThdBoostRatio string `xml:"RachThdBoostRatio"`
	PrachInterfPeriod string `xml:"PrachInterfPeriod"`
	PrachInterfThreshold string `xml:"PrachInterfThreshold"`
	PrachInterfHysteresis string `xml:"PrachInterfHysteresis"`
}

