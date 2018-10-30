package model

import "encoding/xml"

type Iopscfg struct {
	XMLName xml.Name `xml:"IopsCfg"`
	ATTRIBUTES IopscfgAttributes `xml:"attributes"`
}

type IopscfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	IopsSwitch string `xml:"IopsSwitch"`
	EnterIopsThd string `xml:"EnterIopsThd"`
	ExitIopsThd string `xml:"ExitIopsThd"`
	ObjId string `xml:"objId"`
}

