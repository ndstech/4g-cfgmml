package model

import "encoding/xml"

type Distbasedho struct {
	XMLName xml.Name `xml:"DISTBASEDHO"`
	ATTRIBUTES DistbasedhoAttributes `xml:"attributes"`
}

type DistbasedhoAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LocalCellId string `xml:"LocalCellId"`
	DistBasedMeasObjType string `xml:"DistBasedMeasObjType"`
	DistBasedHOThd string `xml:"DistBasedHOThd"`
}

