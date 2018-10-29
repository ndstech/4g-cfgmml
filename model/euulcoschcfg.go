package model

import "encoding/xml"

type Euulcoschcfg struct {
	XMLName xml.Name `xml:"EUULCOSCHCFG"`
	ATTRIBUTES EuulcoschcfgAttributes `xml:"attributes"`
}

type EuulcoschcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	UlInterEnbCamcSw string `xml:"UlInterEnbCamcSw"`
	ObjId string `xml:"objId"`
}

