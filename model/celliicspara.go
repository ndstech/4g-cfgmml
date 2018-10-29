package model

import "encoding/xml"

type Celliicspara struct {
	XMLName xml.Name `xml:"CELLIICSPARA"`
	ATTRIBUTES CelliicsparaAttributes `xml:"attributes"`
}

type CelliicsparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	IicsUserAttrA3Offset string `xml:"IicsUserAttrA3Offset"`
	IicsUserAttrThd string `xml:"IicsUserAttrThd"`
}

