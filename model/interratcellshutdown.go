package model

import "encoding/xml"

type Interratcellshutdown struct {
	XMLName xml.Name `xml:"INTERRATCELLSHUTDOWN"`
	ATTRIBUTES InterratcellshutdownAttributes `xml:"attributes"`
}

type InterratcellshutdownAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	ForceShutdownSwitch string `xml:"ForceShutdownSwitch"`
	StartTime string `xml:"StartTime"`
	StopTime string `xml:"StopTime"`
	BearNumThd string `xml:"BearNumThd"`
	DlPrbThd string `xml:"DlPrbThd"`
	UlPrbThd string `xml:"UlPrbThd"`
	ShutDownType string `xml:"ShutDownType"`
	HighPriArpThd string `xml:"HighPriArpThd"`
	UtranCellDlLoadThd string `xml:"UtranCellDlLoadThd"`
	UtranCellDlLoadOffset string `xml:"UtranCellDlLoadOffset"`
	UtranCellUlLoadThd string `xml:"UtranCellUlLoadThd"`
	UtranCellUlLoadOffset string `xml:"UtranCellUlLoadOffset"`
	GeranCellLoadThd string `xml:"GeranCellLoadThd"`
	GeranCellLoadOffset string `xml:"GeranCellLoadOffset"`
}

