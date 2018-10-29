package model

import "encoding/xml"

type X2interface struct {
	XMLName xml.Name `xml:"X2INTERFACE"`
	ATTRIBUTES X2interfaceAttributes `xml:"attributes"`
}

type X2interfaceAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	X2InterfaceId string `xml:"X2InterfaceId"`
	X2CpBearerId string `xml:"X2CpBearerId"`
	CnOperatorId string `xml:"CnOperatorId"`
	TargetENodeBRelease string `xml:"TargetENodeBRelease"`
	CtrlMode string `xml:"CtrlMode"`
	AutoCfgFlag string `xml:"AutoCfgFlag"`
	CnOpSharingGroupId string `xml:"CnOpSharingGroupId"`
	ObjId string `xml:"objId"`
}

