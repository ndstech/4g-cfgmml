package model

import "encoding/xml"

type Celleabalgopara struct {
	XMLName xml.Name `xml:"CellEABAlgoPara"`
	ATTRIBUTES CelleabalgoparaAttributes `xml:"attributes"`
}

type CelleabalgoparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	EABTriggerThd string `xml:"EABTriggerThd"`
	EABStatPeriod string `xml:"EABStatPeriod"`
	EABCategory string `xml:"EABCategory"`
	EABCancelThd string `xml:"EABCancelThd"`
	EABCancelCondSatiPeriod string `xml:"EABCancelCondSatiPeriod"`
	ABForExceptionData string `xml:"ABForExceptionData"`
	ABForSpecialAC string `xml:"ABForSpecialAC"`
}

