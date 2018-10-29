package model

import "encoding/xml"

type Nodeblicensealmthd struct {
	XMLName xml.Name `xml:"NODEBLICENSEALMTHD"`
	ATTRIBUTES NodeblicensealmthdAttributes `xml:"attributes"`
}

type NodeblicensealmthdAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	OTHD string `xml:"OTHD"`
	OPRD string `xml:"OPRD"`
	RTHD string `xml:"RTHD"`
	RPRD string `xml:"RPRD"`
	OBJID string `xml:"OBJID"`
}

