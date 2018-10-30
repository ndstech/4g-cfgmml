package model

import "encoding/xml"

type Privatecabandcomb struct {
	XMLName xml.Name `xml:"PrivateCaBandComb"`
	ATTRIBUTES PrivatecabandcombAttributes `xml:"attributes"`
}

type PrivatecabandcombAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	PrivateCaCombId string `xml:"PrivateCaCombId"`
	MaxAggregatedBw string `xml:"MaxAggregatedBw"`
	BwCombSetId string `xml:"BwCombSetId"`
	CombBand1Id string `xml:"CombBand1Id"`
	CombBand2Id string `xml:"CombBand2Id"`
	CombBand3Id string `xml:"CombBand3Id"`
	CombBand4Id string `xml:"CombBand4Id"`
	CombBand1Bw string `xml:"CombBand1Bw"`
	CombBand2Bw string `xml:"CombBand2Bw"`
	CombBand3Bw string `xml:"CombBand3Bw"`
	CombBand4Bw string `xml:"CombBand4Bw"`
	CombBand5Id string `xml:"CombBand5Id"`
	CombBand5Bw string `xml:"CombBand5Bw"`
	ObjId string `xml:"objId"`
}

