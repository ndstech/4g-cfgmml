package model

import "encoding/xml"

type Enodebchroutputctrl struct {
	XMLName xml.Name `xml:"ENODEBCHROUTPUTCTRL"`
	ATTRIBUTES EnodebchroutputctrlAttributes `xml:"attributes"`
}

type EnodebchroutputctrlAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	OutPutMode string `xml:"OutPutMode"`
	CallSampleRate string `xml:"CallSampleRate"`
	MaxStoreCall string `xml:"MaxStoreCall"`
	CHRUpLoadingTimeSwitch string `xml:"CHRUpLoadingTimeSwitch"`
	SIGOutputMode string `xml:"SIGOutputMode"`
	ObjId string `xml:"objId"`
}

