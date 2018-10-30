package model

import "encoding/xml"

type Eucoschcfg struct {
	XMLName xml.Name `xml:"EuCoSchCfg"`
	ATTRIBUTES EucoschcfgAttributes `xml:"attributes"`
}

type EucoschcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	PrtNodeBaseBandEqmId string `xml:"PrtNodeBaseBandEqmId"`
	SchNodeBaseBandEqmId string `xml:"SchNodeBaseBandEqmId"`
	WorkMode string `xml:"WorkMode"`
	ObjId string `xml:"objId"`
}

