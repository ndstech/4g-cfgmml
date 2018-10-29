package model

import "encoding/xml"

type Cnoperatorstandardqci struct {
	XMLName xml.Name `xml:"CNOPERATORSTANDARDQCI"`
	ATTRIBUTES CnoperatorstandardqciAttributes `xml:"attributes"`
}

type CnoperatorstandardqciAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CnOperatorId string `xml:"CnOperatorId"`
	Qci string `xml:"Qci"`
	ServiceIrHoCfgGroupId string `xml:"ServiceIrHoCfgGroupId"`
	ServiceIfHoCfgGroupId string `xml:"ServiceIfHoCfgGroupId"`
	ServiceHoBearerPolicy string `xml:"ServiceHoBearerPolicy"`
	LocalQciArp string `xml:"LocalQciArp"`
}

