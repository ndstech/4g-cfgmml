package model

import "encoding/xml"

type Enodebintegritycap struct {
	XMLName xml.Name `xml:"ENodeBIntegrityCap"`
	ATTRIBUTES EnodebintegritycapAttributes `xml:"attributes"`
}

type EnodebintegritycapAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	PrimaryIntegrityAlgo string `xml:"PrimaryIntegrityAlgo"`
	SecondIntegrityAlgo string `xml:"SecondIntegrityAlgo"`
	ThirdIntegrityAlgo string `xml:"ThirdIntegrityAlgo"`
	NullAlgo string `xml:"NullAlgo"`
	ObjId string `xml:"objId"`
}

