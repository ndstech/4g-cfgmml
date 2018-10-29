package model

import "encoding/xml"

type Mmefeaturecfg struct {
	XMLName xml.Name `xml:"MMEFEATURECFG"`
	ATTRIBUTES MmefeaturecfgAttributes `xml:"attributes"`
}

type MmefeaturecfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	S1InterfaceId string `xml:"S1InterfaceId"`
	MdtEnable string `xml:"MdtEnable"`
	CtrlMode string `xml:"CtrlMode"`
}

