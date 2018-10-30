package model

import "encoding/xml"

type Mbmspara struct {
	XMLName xml.Name `xml:"MBMSPara"`
	ATTRIBUTES MbmsparaAttributes `xml:"attributes"`
}

type MbmsparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	SyncPeriod string `xml:"SyncPeriod"`
	MbmsCtrlSwitch string `xml:"MbmsCtrlSwitch"`
	M2DisconnProtectTime string `xml:"M2DisconnProtectTime"`
	ObjId string `xml:"objId"`
}

