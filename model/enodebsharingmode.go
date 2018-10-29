package model

import "encoding/xml"

type Enodebsharingmode struct {
	XMLName xml.Name `xml:"ENODEBSHARINGMODE"`
	ATTRIBUTES EnodebsharingmodeAttributes `xml:"attributes"`
}

type EnodebsharingmodeAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	ENodeBSharingMode string `xml:"ENodeBSharingMode"`
	ObjId string `xml:"objId"`
}

