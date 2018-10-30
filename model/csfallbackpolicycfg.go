package model

import "encoding/xml"

type Csfallbackpolicycfg struct {
	XMLName xml.Name `xml:"CSFallBackPolicyCfg"`
	ATTRIBUTES CsfallbackpolicycfgAttributes `xml:"attributes"`
}

type CsfallbackpolicycfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CsfbHoPolicyCfg string `xml:"CsfbHoPolicyCfg"`
	IdleModeCsfbHoPolicyCfg string `xml:"IdleModeCsfbHoPolicyCfg"`
	CsfbUserArpCfgSwitch string `xml:"CsfbUserArpCfgSwitch"`
	NormalCsfbUserArp string `xml:"NormalCsfbUserArp"`
	ObjId string `xml:"objId"`
}

