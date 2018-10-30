package model

import "encoding/xml"

type Cellmlbautogroup struct {
	XMLName xml.Name `xml:"CellMlbAutoGroup"`
	ATTRIBUTES CellmlbautogroupAttributes `xml:"attributes"`
}

type CellmlbautogroupAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	OverlapIndAddThd string `xml:"OverlapIndAddThd"`
	OverlapIndDelThd string `xml:"OverlapIndDelThd"`
	StatPeriod string `xml:"StatPeriod"`
	SleepPeriod string `xml:"SleepPeriod"`
	AddedMeasCfg string `xml:"AddedMeasCfg"`
	OverlapIndAutoSampleNum string `xml:"OverlapIndAutoSampleNum"`
	OptMode string `xml:"OptMode"`
	MicroOverlapIndAddThd string `xml:"MicroOverlapIndAddThd"`
	MicroOverlapIndDelThd string `xml:"MicroOverlapIndDelThd"`
}

