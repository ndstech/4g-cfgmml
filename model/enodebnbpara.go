package model

import "encoding/xml"

type Enodebnbpara struct {
	XMLName xml.Name `xml:"ENodeBNbPara"`
	ATTRIBUTES EnodebnbparaAttributes `xml:"attributes"`
}

type EnodebnbparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	NbRsvMinUserNumRatio string `xml:"NbRsvMinUserNumRatio"`
	ObjId string `xml:"objId"`
}

