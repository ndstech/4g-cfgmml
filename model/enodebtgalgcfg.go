package model

import "encoding/xml"

type Enodebtgalgcfg struct {
	XMLName xml.Name `xml:"eNodeBTGAlgCfg"`
	ATTRIBUTES EnodebtgalgcfgAttributes `xml:"attributes"`
}

type EnodebtgalgcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	FilterCoeffTG string `xml:"FilterCoeffTG"`
	StatPeriodTG string `xml:"StatPeriodTG"`
	ObjId string `xml:"objId"`
}

