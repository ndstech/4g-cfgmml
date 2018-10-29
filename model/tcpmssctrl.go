package model

import "encoding/xml"

type Tcpmssctrl struct {
	XMLName xml.Name `xml:"TCPMSSCTRL"`
	ATTRIBUTES TcpmssctrlAttributes `xml:"attributes"`
}

type TcpmssctrlAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	TcpMssCtrlSwitch string `xml:"TcpMssCtrlSwitch"`
	TcpMssThd string `xml:"TcpMssThd"`
	ObjId string `xml:"objId"`
}

