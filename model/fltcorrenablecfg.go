package model

import "encoding/xml"

type Fltcorrenablecfg struct {
	XMLName xml.Name `xml:"FLTCORRENABLECFG"`
	ATTRIBUTES FltcorrenablecfgAttributes `xml:"attributes"`
}

type FltcorrenablecfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	ENABLERELA string `xml:"ENABLERELA"`
}

