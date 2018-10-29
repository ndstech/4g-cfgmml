package model

import "encoding/xml"

type Gbtsglobalpara struct {
	XMLName xml.Name `xml:"GBTSGLOBALPARA"`
	ATTRIBUTES GbtsglobalparaAttributes `xml:"attributes"`
}

type GbtsglobalparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	OBJID string `xml:"OBJID"`
	MCSTANDARD string `xml:"MCSTANDARD"`
}

