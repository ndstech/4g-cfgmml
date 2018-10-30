package model

import "encoding/xml"

type Cellshutdown struct {
	XMLName xml.Name `xml:"CellShutdown"`
	ATTRIBUTES CellshutdownAttributes `xml:"attributes"`
}

type CellshutdownAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	CellShutdownSwitch string `xml:"CellShutdownSwitch"`
	StartTime string `xml:"StartTime"`
	StopTime string `xml:"StopTime"`
	DlPrbThd string `xml:"DlPrbThd"`
	DlPrbOffset string `xml:"DlPrbOffset"`
	UlPrbThd string `xml:"UlPrbThd"`
	UlPrbOffset string `xml:"UlPrbOffset"`
	ForceShutdownUENumThd string `xml:"ForceShutdownUENumThd"`
	PunishTime string `xml:"PunishTime"`
}

