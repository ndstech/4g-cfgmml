package model

import "encoding/xml"

type Interratpolicycfggroup struct {
	XMLName xml.Name `xml:"INTERRATPOLICYCFGGROUP"`
	ATTRIBUTES InterratpolicycfggroupAttributes `xml:"attributes"`
}

type InterratpolicycfggroupAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	InterRatPolicyCfgGroupId string `xml:"InterRatPolicyCfgGroupId"`
	UtranHoCfg string `xml:"UtranHoCfg"`
	GeranGsmHoCfg string `xml:"GeranGsmHoCfg"`
	GeranGprsEdgeHoCfg string `xml:"GeranGprsEdgeHoCfg"`
	Cdma1xRttHoCfg string `xml:"Cdma1xRttHoCfg"`
	CdmaHrpdHoCfg string `xml:"CdmaHrpdHoCfg"`
	NoHoFlag string `xml:"NoHoFlag"`
	NoFastAnrFlag string `xml:"NoFastAnrFlag"`
	ObjId string `xml:"objId"`
}

