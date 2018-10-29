package model

import "encoding/xml"

type Celllowpower struct {
	XMLName xml.Name `xml:"CELLLOWPOWER"`
	ATTRIBUTES CelllowpowerAttributes `xml:"attributes"`
}

type CelllowpowerAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	LowPwrSwitch string `xml:"LowPwrSwitch"`
	CellUsedPwrReduceTimeLen string `xml:"CellUsedPwrReduceTimeLen"`
	RsPwrReduceTimeLen string `xml:"RsPwrReduceTimeLen"`
	RfShutDownTimeLen string `xml:"RfShutDownTimeLen"`
	CellUsedPwrRatio string `xml:"CellUsedPwrRatio"`
	RsPwrAdjOffset string `xml:"RsPwrAdjOffset"`
	EnterTimeLen string `xml:"EnterTimeLen"`
	BakPwrSavPolicy string `xml:"BakPwrSavPolicy"`
}

