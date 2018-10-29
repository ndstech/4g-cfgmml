package model

import "encoding/xml"

type Enodebfunction struct {
	XMLName xml.Name `xml:"ENODEBFUNCTION"`
	ATTRIBUTES EnodebfunctionAttributes `xml:"attributes"`
}

type EnodebfunctionAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	ObjId string `xml:"objId"`
	ENodeBFunctionName string `xml:"eNodeBFunctionName"`
	ApplicationRef string `xml:"ApplicationRef"`
	ENodeBId string `xml:"eNodeBId"`
	UserLabel string `xml:"UserLabel"`
	NermVersion string `xml:"NermVersion"`
	ProductVersion string `xml:"ProductVersion"`
	ProductInterfaceID string `xml:"ProductInterfaceID"`
	LmtVersion string `xml:"LmtVersion"`
}

