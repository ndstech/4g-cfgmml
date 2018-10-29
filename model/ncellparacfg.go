package model

import "encoding/xml"

type Ncellparacfg struct {
	XMLName xml.Name `xml:"NCELLPARACFG"`
	ATTRIBUTES NcellparacfgAttributes `xml:"attributes"`
}

type NcellparacfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	RatType string `xml:"RatType"`
	NCellOdDisThd string `xml:"NCellOdDisThd"`
	HoStatThd string `xml:"HoStatThd"`
	HoSuccThd string `xml:"HoSuccThd"`
	NcellNumForAnr string `xml:"NcellNumForAnr"`
}

