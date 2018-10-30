package model

import "encoding/xml"

type Eucellalmcfg struct {
	XMLName xml.Name `xml:"EuCellAlmCfg"`
	ATTRIBUTES EucellalmcfgAttributes `xml:"attributes"`
}

type EucellalmcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	RxNoisePwrThd string `xml:"RxNoisePwrThd"`
}

