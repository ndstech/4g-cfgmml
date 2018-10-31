package model

import "encoding/xml"

type Cagroupscellcfg struct {
	XMLName    xml.Name                  `xml:"CaGroupSCellCfg"`
	ATTRIBUTES CagroupscellcfgAttributes `xml:"attributes"`
}

type CagroupscellcfgAttributes struct {
	XMLName           xml.Name `xml:"attributes"`
	LocalCellId       string   `xml:"LocalCellId"`
	SCellENODEB_ID    string   `xml:"SCelleNodeBId"`
	SCellLocalCellId  string   `xml:"SCellLocalCellId"`
	SCellBlindCfgFlag string   `xml:"SCellBlindCfgFlag"`
	SCellPriority     string   `xml:"SCellPriority"`
	SCellA4Offset     string   `xml:"SCellA4Offset"`
	SCellA2Offset     string   `xml:"SCellA2Offset"`
	SpidGrpId         string   `xml:"SpidGrpId"`
}
