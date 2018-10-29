package model

import "encoding/xml"

type Asparagroup struct {
	XMLName xml.Name `xml:"ASPARAGROUP"`
	ATTRIBUTES AsparagroupAttributes `xml:"attributes"`
}

type AsparagroupAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	AsParaGroupID string `xml:"AsParaGroupID"`
	AsPreallocDuration string `xml:"AsPreallocDuration"`
	AsPreallocMinPeriod string `xml:"AsPreallocMinPeriod"`
	AsPreallocSize string `xml:"AsPreallocSize"`
	AsSchPriFactor string `xml:"AsSchPriFactor"`
	ObjId string `xml:"objId"`
}

