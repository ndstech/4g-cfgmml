package model

import "encoding/xml"

type Nbcellulschcealgo struct {
	XMLName xml.Name `xml:"NBCELLULSCHCEALGO"`
	ATTRIBUTES NbcellulschcealgoAttributes `xml:"attributes"`
}

type NbcellulschcealgoAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	CoverageLevel string `xml:"CoverageLevel"`
	UlInitialMcs string `xml:"UlInitialMcs"`
	UlInitialTransRptCount string `xml:"UlInitialTransRptCount"`
	AckNackTransRptCount string `xml:"AckNackTransRptCount"`
	AckNackTransRptCountMsg4 string `xml:"AckNackTransRptCountMsg4"`
	NbLogicChSrProhibitTimer string `xml:"NbLogicChSrProhibitTimer"`
}

