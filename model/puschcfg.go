package model

import "encoding/xml"

type Puschcfg struct {
	XMLName xml.Name `xml:"PUSCHCFG"`
	ATTRIBUTES PuschcfgAttributes `xml:"attributes"`
}

type PuschcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	Nsb string `xml:"Nsb"`
	HoppingMode string `xml:"HoppingMode"`
	HoppingOffset string `xml:"HoppingOffset"`
	GroupHoppingEnabled string `xml:"GroupHoppingEnabled"`
	GroupAssignPUSCH string `xml:"GroupAssignPUSCH"`
	SeqHoppingEnabled string `xml:"SeqHoppingEnabled"`
	CyclicShift string `xml:"CyclicShift"`
	Qam64Enabled string `xml:"Qam64Enabled"`
	R12Qam64Enabled string `xml:"R12Qam64Enabled"`
}

