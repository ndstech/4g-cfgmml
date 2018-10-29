package model

import "encoding/xml"

type Enblicensealmthd struct {
	XMLName xml.Name `xml:"ENBLICENSEALMTHD"`
	ATTRIBUTES EnblicensealmthdAttributes `xml:"attributes"`
}

type EnblicensealmthdAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	OPRD string `xml:"OPRD"`
	OTHD string `xml:"OTHD"`
	RPRD string `xml:"RPRD"`
	RTHD string `xml:"RTHD"`
	ObjId string `xml:"objId"`
}

