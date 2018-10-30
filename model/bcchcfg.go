package model

import "encoding/xml"

type Bcchcfg struct {
	XMLName xml.Name `xml:"BcchCfg"`
	ATTRIBUTES BcchcfgAttributes `xml:"attributes"`
}

type BcchcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	ModifyPeriodCoeff string `xml:"ModifyPeriodCoeff"`
	ModifyPeriodCoeffForNb string `xml:"ModifyPeriodCoeffForNb"`
}

