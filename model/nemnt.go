package model

import "encoding/xml"

type Nemnt struct {
	XMLName xml.Name `xml:"NEMNT"`
	ATTRIBUTES NemntAttributes `xml:"attributes"`
}

type NemntAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	MNTMODE string `xml:"MNTMODE"`
	ST string `xml:"ST"`
	ET string `xml:"ET"`
	MMSETREMARK string `xml:"MMSETREMARK"`
}

