package model

import "encoding/xml"

type Srsadaptivecfg struct {
	XMLName xml.Name `xml:"SRSADAPTIVECFG"`
	ATTRIBUTES SrsadaptivecfgAttributes `xml:"attributes"`
}

type SrsadaptivecfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	SrsPeriodAdaptive string `xml:"SrsPeriodAdaptive"`
	ObjId string `xml:"objId"`
}

