package model

import "encoding/xml"

type Ncellclassmgt struct {
	XMLName xml.Name `xml:"NCELLCLASSMGT"`
	ATTRIBUTES NcellclassmgtAttributes `xml:"attributes"`
}

type NcellclassmgtAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	StatPeriodForNcellClass string `xml:"StatPeriodForNcellClass"`
	HoAttemptThd string `xml:"HoAttemptThd"`
	CaSCellCfgThd string `xml:"CaSCellCfgThd"`
	HoSuccThd string `xml:"HoSuccThd"`
	IntraRatNcellMgtMode string `xml:"IntraRatNcellMgtMode"`
	ObjId string `xml:"objId"`
}

