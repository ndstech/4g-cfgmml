package model

import "encoding/xml"

type Gerannfreqgrouparfcn struct {
	XMLName xml.Name `xml:"GERANNFREQGROUPARFCN"`
	ATTRIBUTES GerannfreqgrouparfcnAttributes `xml:"attributes"`
}

type GerannfreqgrouparfcnAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	BcchGroupId string `xml:"BcchGroupId"`
	GeranArfcn string `xml:"GeranArfcn"`
}

