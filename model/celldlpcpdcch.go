package model

import "encoding/xml"

type Celldlpcpdcch struct {
	XMLName xml.Name `xml:"CELLDLPCPDCCH"`
	ATTRIBUTES CelldlpcpdcchAttributes `xml:"attributes"`
}

type CelldlpcpdcchAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	DediDciPwrOffset string `xml:"DediDciPwrOffset"`
}

