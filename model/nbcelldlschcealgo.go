package model

import "encoding/xml"

type Nbcelldlschcealgo struct {
	XMLName xml.Name `xml:"NbCellDlSchCEAlgo"`
	ATTRIBUTES NbcelldlschcealgoAttributes `xml:"attributes"`
}

type NbcelldlschcealgoAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	CoverageLevel string `xml:"CoverageLevel"`
	DlInitialTransRptCount string `xml:"DlInitialTransRptCount"`
	DlInitialMcs string `xml:"DlInitialMcs"`
	UuMessageWaitingTimer string `xml:"UuMessageWaitingTimer"`
}

