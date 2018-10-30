package model

import "encoding/xml"

type Cellrfshutdown struct {
	XMLName xml.Name `xml:"CellRfShutdown"`
	ATTRIBUTES CellrfshutdownAttributes `xml:"attributes"`
}

type CellrfshutdownAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	RfShutdownSwitch string `xml:"RfShutdownSwitch"`
	StartTime string `xml:"StartTime"`
	StopTime string `xml:"StopTime"`
	RsPwrAdjOffset string `xml:"RsPwrAdjOffset"`
	DlPrbThd string `xml:"DlPrbThd"`
	UlPrbThd string `xml:"UlPrbThd"`
	DlPrbOffset string `xml:"DlPrbOffset"`
	UlPrbOffset string `xml:"UlPrbOffset"`
	UENumThd string `xml:"UENumThd"`
	RfShutdownJudgingPolicy string `xml:"RfShutdownJudgingPolicy"`
}

