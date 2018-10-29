package model

import "encoding/xml"

type Phichcfg struct {
	XMLName xml.Name `xml:"PHICHCFG"`
	ATTRIBUTES PhichcfgAttributes `xml:"attributes"`
}

type PhichcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	PhichDuration string `xml:"PhichDuration"`
	PhichResource string `xml:"PhichResource"`
}

