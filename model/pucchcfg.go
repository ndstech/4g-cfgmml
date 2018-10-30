package model

import "encoding/xml"

type Pucchcfg struct {
	XMLName xml.Name `xml:"PUCCHCfg"`
	ATTRIBUTES PucchcfgAttributes `xml:"attributes"`
}

type PucchcfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	DeltaShift string `xml:"DeltaShift"`
	NaSriChNum string `xml:"NaSriChNum"`
	CqiRbNum string `xml:"CqiRbNum"`
	PucchExtendedRBNum string `xml:"PucchExtendedRBNum"`
	Format1ChAllocMode string `xml:"Format1ChAllocMode"`
	SriPeriodAdaptive string `xml:"SriPeriodAdaptive"`
	Format2ChAllocMode string `xml:"Format2ChAllocMode"`
	PucchAllocPolicy string `xml:"PucchAllocPolicy"`
	Format3RbNum string `xml:"Format3RbNum"`
	Max2CCAckChNum string `xml:"Max2CCAckChNum"`
}

