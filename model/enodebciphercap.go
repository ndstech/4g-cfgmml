package model

import "encoding/xml"

type Enodebciphercap struct {
	XMLName xml.Name `xml:"ENodeBCipherCap"`
	ATTRIBUTES EnodebciphercapAttributes `xml:"attributes"`
}

type EnodebciphercapAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	PrimaryCipherAlgo string `xml:"PrimaryCipherAlgo"`
	SecondCipherAlgo string `xml:"SecondCipherAlgo"`
	ThirdCipherAlgo string `xml:"ThirdCipherAlgo"`
	FourthCipherAlgo string `xml:"FourthCipherAlgo"`
	ObjId string `xml:"objId"`
}

