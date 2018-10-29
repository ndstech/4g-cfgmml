package model

import "encoding/xml"

type Paraautooptcfg struct {
	XMLName xml.Name `xml:"PARAAUTOOPTCFG"`
	ATTRIBUTES ParaautooptcfgAttributes `xml:"attributes"`
}

type ParaautooptcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	PUSCHRsrpHighThd4AutoOpt string `xml:"PUSCHRsrpHighThd4AutoOpt"`
	PUCCHPcSINROffset4AutoOpt string `xml:"PUCCHPcSINROffset4AutoOpt"`
	P0NominalPUCCH4AutoOpt string `xml:"P0NominalPUCCH4AutoOpt"`
	HOTimesThd string `xml:"HOTimesThd"`
	ObjId string `xml:"objId"`
}

