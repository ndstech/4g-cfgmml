package model

import "encoding/xml"

type Rrujointcalparacfg struct {
	XMLName xml.Name `xml:"RRUJOINTCALPARACFG"`
	ATTRIBUTES RrujointcalparacfgAttributes `xml:"attributes"`
}

type RrujointcalparacfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	TxChnCalSwitch string `xml:"TxChnCalSwitch"`
	TxChnCalTime string `xml:"TxChnCalTime"`
	TxChnCalPeriod string `xml:"TxChnCalPeriod"`
	TxChnAntSpacing string `xml:"TxChnAntSpacing"`
	AauPassivePortCalibPeriod string `xml:"AauPassivePortCalibPeriod"`
}

