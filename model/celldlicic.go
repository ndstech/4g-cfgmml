package model

import "encoding/xml"

type Celldlicic struct {
	XMLName xml.Name `xml:"CELLDLICIC"`
	ATTRIBUTES CelldlicicAttributes `xml:"attributes"`
}

type CelldlicicAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	BandMode string `xml:"BandMode"`
	DlIcicUserAttrGfactorThd string `xml:"DlIcicUserAttrGfactorThd"`
	DlIcicNoiseUserRsrpThd string `xml:"DlIcicNoiseUserRsrpThd"`
	AIcIcPlusA3Offset string `xml:"AIcIcPlusA3Offset"`
	AIcicPlusPCAdjRange string `xml:"AIcicPlusPCAdjRange"`
}

