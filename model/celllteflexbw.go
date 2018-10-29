package model

import "encoding/xml"

type Celllteflexbw struct {
	XMLName xml.Name `xml:"CELLLTEFLEXBW"`
	ATTRIBUTES CelllteflexbwAttributes `xml:"attributes"`
}

type CelllteflexbwAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	LteFlexBwSwitch string `xml:"LteFlexBwSwitch"`
	DlCustStartPrbIndex string `xml:"DlCustStartPrbIndex"`
	DlCustEndPrbIndex string `xml:"DlCustEndPrbIndex"`
	UlCustStartPrbIndex string `xml:"UlCustStartPrbIndex"`
	UlCustEndPrbIndex string `xml:"UlCustEndPrbIndex"`
}

