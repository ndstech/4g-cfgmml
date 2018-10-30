package model

import "encoding/xml"

type Cagroupcell struct {
	XMLName xml.Name `xml:"CaGroupCell"`
	ATTRIBUTES CagroupcellAttributes `xml:"attributes"`
}

type CagroupcellAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CaGroupId string `xml:"CaGroupId"`
	LocalCellId string `xml:"LocalCellId"`
	ENodeBId string `xml:"eNodeBId"`
	PreferredPCellPriority string `xml:"PreferredPCellPriority"`
	PCellA4RsrpThd string `xml:"PCellA4RsrpThd"`
	PCellA4RsrqThd string `xml:"PCellA4RsrqThd"`
}

