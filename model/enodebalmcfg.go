package model

import "encoding/xml"

type Enodebalmcfg struct {
	XMLName xml.Name `xml:"ENODEBALMCFG"`
	ATTRIBUTES EnodebalmcfgAttributes `xml:"attributes"`
}

type EnodebalmcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	TrafficFaultyDetectPeriod string `xml:"TrafficFaultyDetectPeriod"`
	SrvIntfAlmCfgSw string `xml:"SrvIntfAlmCfgSw"`
	ObjId string `xml:"objId"`
}

