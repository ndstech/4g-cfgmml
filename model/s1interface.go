package model

import "encoding/xml"

type S1interface struct {
	XMLName xml.Name `xml:"S1INTERFACE"`
	ATTRIBUTES S1interfaceAttributes `xml:"attributes"`
}

type S1interfaceAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	S1InterfaceId string `xml:"S1InterfaceId"`
	S1CpBearerId string `xml:"S1CpBearerId"`
	CnOperatorId string `xml:"CnOperatorId"`
	MmeRelease string `xml:"MmeRelease"`
	S1InterfaceIsBlock string `xml:"S1InterfaceIsBlock"`
	CtrlMode string `xml:"CtrlMode"`
	AutoCfgFlag string `xml:"AutoCfgFlag"`
	Priority string `xml:"Priority"`
	CnOpSharingGroupId string `xml:"CnOpSharingGroupId"`
	ServedGummeis string `xml:"ServedGummeis"`
	ObjId string `xml:"objId"`
}

