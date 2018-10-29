package model

import "encoding/xml"

type Geraninterfcfg struct {
	XMLName xml.Name `xml:"GERANINTERFCFG"`
	ATTRIBUTES GeraninterfcfgAttributes `xml:"attributes"`
}

type GeraninterfcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	DlGeranIntefRbNum string `xml:"DlGeranIntefRbNum"`
	DlRbAvailSinrThd string `xml:"DlRbAvailSinrThd"`
}

