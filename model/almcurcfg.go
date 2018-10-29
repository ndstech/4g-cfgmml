package model

import "encoding/xml"

type Almcurcfg struct {
	XMLName xml.Name `xml:"ALMCURCFG"`
	ATTRIBUTES AlmcurcfgAttributes `xml:"attributes"`
}

type AlmcurcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	AID string `xml:"AID"`
	ALVL string `xml:"ALVL"`
	ASS string `xml:"ASS"`
	SHLDFLG string `xml:"SHLDFLG"`
	ANM string `xml:"ANM"`
}

