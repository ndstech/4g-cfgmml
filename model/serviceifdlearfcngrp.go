package model

import "encoding/xml"

type Serviceifdlearfcngrp struct {
	XMLName xml.Name `xml:"ServiceIfDlEarfcnGrp"`
	ATTRIBUTES ServiceifdlearfcngrpAttributes `xml:"attributes"`
}

type ServiceifdlearfcngrpAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CnOperatorId string `xml:"CnOperatorId"`
	ServiceIfHoCfgGroupId string `xml:"ServiceIfHoCfgGroupId"`
	DlEarfcnIndex string `xml:"DlEarfcnIndex"`
	DlEarfcn string `xml:"DlEarfcn"`
	ServiceHoFreqPriority string `xml:"ServiceHoFreqPriority"`
}

