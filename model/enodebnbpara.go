package model

import "encoding/xml"

type Enodebnbpara struct {
	XMLName xml.Name `xml:"ENODEBNBPARA"`
	ATTRIBUTES EnodebnbparaAttributes `xml:"attributes"`
}

type EnodebnbparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	NbRsvMinUserNumRatio string `xml:"NbRsvMinUserNumRatio"`
	ObjId string `xml:"objId"`
}

