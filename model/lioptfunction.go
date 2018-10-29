package model

import "encoding/xml"

type Lioptfunction struct {
	XMLName xml.Name `xml:"LIOPTFUNCTION"`
	ATTRIBUTES LioptfunctionAttributes `xml:"attributes"`
}

type LioptfunctionAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	IOptFunctionID string `xml:"IOptFunctionID"`
	IOptFunctionName string `xml:"IOptFunctionName"`
	IOptFeatureID string `xml:"IOptFeatureID"`
	MeasureObjType string `xml:"MeasureObjType"`
	ObjId string `xml:"objId"`
}

