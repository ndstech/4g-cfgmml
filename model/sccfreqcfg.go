package model

import "encoding/xml"

type Sccfreqcfg struct {
	XMLName xml.Name `xml:"SccFreqCfg"`
	ATTRIBUTES SccfreqcfgAttributes `xml:"attributes"`
}

type SccfreqcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	PccDlEarfcn string `xml:"PccDlEarfcn"`
	SccDlEarfcn string `xml:"SccDlEarfcn"`
	SccPriority string `xml:"SccPriority"`
	SccA2Offset string `xml:"SccA2Offset"`
	SccA4Offset string `xml:"SccA4Offset"`
	BlindScellAddThd string `xml:"BlindScellAddThd"`
	BlindScellDelThd string `xml:"BlindScellDelThd"`
	CtrlMode string `xml:"CtrlMode"`
	SpidGrpId string `xml:"SpidGrpId"`
}

