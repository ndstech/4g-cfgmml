package model

import "encoding/xml"

type Pccfreqcfg struct {
	XMLName xml.Name `xml:"PCCFREQCFG"`
	ATTRIBUTES PccfreqcfgAttributes `xml:"attributes"`
}

type PccfreqcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	PccDlEarfcn string `xml:"PccDlEarfcn"`
	PreferredPccPriority string `xml:"PreferredPccPriority"`
	PccA4RsrpThd string `xml:"PccA4RsrpThd"`
	PccA4RsrqThd string `xml:"PccA4RsrqThd"`
	ObjId string `xml:"objId"`
}

