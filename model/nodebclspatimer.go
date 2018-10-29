package model

import "encoding/xml"

type Nodebclspatimer struct {
	XMLName xml.Name `xml:"NODEBCLSPATIMER"`
	ATTRIBUTES NodebclspatimerAttributes `xml:"attributes"`
}

type NodebclspatimerAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LOWERLIMIT string `xml:"LOWERLIMIT"`
	UPPERLIMIT string `xml:"UPPERLIMIT"`
	OBJID string `xml:"OBJID"`
}

