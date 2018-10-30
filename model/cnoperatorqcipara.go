package model

import "encoding/xml"

type Cnoperatorqcipara struct {
	XMLName xml.Name `xml:"CnOperatorQciPara"`
	ATTRIBUTES CnoperatorqciparaAttributes `xml:"attributes"`
}

type CnoperatorqciparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CnOperatorId string `xml:"CnOperatorId"`
	Qci string `xml:"Qci"`
	ServiceIrHoCfgGroupId string `xml:"ServiceIrHoCfgGroupId"`
	ServiceIfHoCfgGroupId string `xml:"ServiceIfHoCfgGroupId"`
	ServiceHoBearerPolicy string `xml:"ServiceHoBearerPolicy"`
	LocalQciArp string `xml:"LocalQciArp"`
}

