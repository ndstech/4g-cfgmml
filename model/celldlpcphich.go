package model

import "encoding/xml"

type Celldlpcphich struct {
	XMLName xml.Name `xml:"CELLDLPCPHICH"`
	ATTRIBUTES CelldlpcphichAttributes `xml:"attributes"`
}

type CelldlpcphichAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	PwrOffset string `xml:"PwrOffset"`
}

