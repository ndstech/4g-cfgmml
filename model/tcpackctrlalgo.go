package model

import "encoding/xml"

type Tcpackctrlalgo struct {
	XMLName xml.Name `xml:"TCPACKCTRLALGO"`
	ATTRIBUTES TcpackctrlalgoAttributes `xml:"attributes"`
}

type TcpackctrlalgoAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	AckCtrlSwitch string `xml:"AckCtrlSwitch"`
	DlMaxThroughput string `xml:"DlMaxThroughput"`
	CtrlTimerLength string `xml:"CtrlTimerLength"`
	ObjId string `xml:"objId"`
}

