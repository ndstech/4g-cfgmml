package model

import "encoding/xml"

type Srscfg struct {
	XMLName xml.Name `xml:"SRSCFG"`
	ATTRIBUTES SrscfgAttributes `xml:"attributes"`
}

type SrscfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	SrsSubframeCfg string `xml:"SrsSubframeCfg"`
	AnSrsSimuTrans string `xml:"AnSrsSimuTrans"`
	SrsCfgInd string `xml:"SrsCfgInd"`
	TddSrsCfgMode string `xml:"TddSrsCfgMode"`
	FddSrsCfgMode string `xml:"FddSrsCfgMode"`
	SrsAlgoOptSwitch string `xml:"SrsAlgoOptSwitch"`
	SrsCfgPolicySwitch string `xml:"SrsCfgPolicySwitch"`
	SrsResOptSwitch string `xml:"SrsResOptSwitch"`
}

