package model

import "encoding/xml"

type X2 struct {
	XMLName xml.Name `xml:"X2"`
	ATTRIBUTES X2Attributes `xml:"attributes"`
}

type X2Attributes struct {
	XMLName xml.Name `xml:"attributes"`
	X2Id string `xml:"X2Id"`
	CnOperatorId string `xml:"CnOperatorId"`
	CpEpGroupId string `xml:"CpEpGroupId"`
	UpEpGroupId string `xml:"UpEpGroupId"`
	TargetENodeBRelease string `xml:"TargetENodeBRelease"`
	UserLabel string `xml:"UserLabel"`
	EpGroupCfgFlag string `xml:"EpGroupCfgFlag"`
	CnOpSharingGroupId string `xml:"CnOpSharingGroupId"`
	ObjId string `xml:"objId"`
}

