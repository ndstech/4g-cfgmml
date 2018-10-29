package model

import "encoding/xml"

type Enodebsondbcfg struct {
	XMLName xml.Name `xml:"ENODEBSONDBCFG"`
	ATTRIBUTES EnodebsondbcfgAttributes `xml:"attributes"`
}

type EnodebsondbcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	StartTime string `xml:"StartTime"`
	ObjId string `xml:"objId"`
}

