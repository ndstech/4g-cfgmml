package model

import "encoding/xml"

type Nodebsmthpwr struct {
	XMLName xml.Name `xml:"NODEBSMTHPWR"`
	ATTRIBUTES NodebsmthpwrAttributes `xml:"attributes"`
}

type NodebsmthpwrAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	SMTHPWRSWITCH string `xml:"SMTHPWRSWITCH"`
	STEP string `xml:"STEP"`
	PERIOD string `xml:"PERIOD"`
	OBJID string `xml:"OBJID"`
}

