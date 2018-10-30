package model

import "encoding/xml"

type Enodebsharingmode struct {
	XMLName xml.Name `xml:"ENodeBSharingMode"`
	ATTRIBUTES EnodebsharingmodeAttributes `xml:"attributes"`
}

type EnodebsharingmodeAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	ENodeBSharingMode string `xml:"ENodeBSharingMode"`
	ObjId string `xml:"objId"`
}

