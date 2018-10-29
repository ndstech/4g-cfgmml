package model

import "encoding/xml"

type S1 struct {
	XMLName xml.Name `xml:"S1"`
	ATTRIBUTES S1Attributes `xml:"attributes"`
}

type S1Attributes struct {
	XMLName xml.Name `xml:"attributes"`
	S1Id string `xml:"S1Id"`
	CnOperatorId string `xml:"CnOperatorId"`
	UpEpGroupId string `xml:"UpEpGroupId"`
	MmeRelease string `xml:"MmeRelease"`
	UserLabel string `xml:"UserLabel"`
	EpGroupCfgFlag string `xml:"EpGroupCfgFlag"`
	Priority string `xml:"Priority"`
	CnOpSharingGroupId string `xml:"CnOpSharingGroupId"`
	MdtEnable string `xml:"MdtEnable"`
	ObjId string `xml:"objId"`
}

