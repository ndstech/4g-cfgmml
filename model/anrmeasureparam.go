package model

import "encoding/xml"

type Anrmeasureparam struct {
	XMLName xml.Name `xml:"ANRMeasureParam"`
	ATTRIBUTES AnrmeasureparamAttributes `xml:"attributes"`
}

type AnrmeasureparamAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	A3TimeToTrigForANR string `xml:"A3TimeToTrigForANR"`
	A3HystForANR string `xml:"A3HystForANR"`
	A3OffsetForANR string `xml:"A3OffsetForANR"`
	A4TimeToTrigForANR string `xml:"A4TimeToTrigForANR"`
	A4HystForANR string `xml:"A4HystForANR"`
	A4ThdRsrpForANR string `xml:"A4ThdRsrpForANR"`
	ObjId string `xml:"objId"`
}

