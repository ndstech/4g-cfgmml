package model

import "encoding/xml"

type Nodebresallocrule struct {
	XMLName xml.Name `xml:"NODEBRESALLOCRULE"`
	ATTRIBUTES NodebresallocruleAttributes `xml:"attributes"`
}

type NodebresallocruleAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	RULE string `xml:"RULE"`
	SW string `xml:"SW"`
	CSUSERNUM string `xml:"CSUSERNUM"`
	PSUSERNUM string `xml:"PSUSERNUM"`
	OBJID string `xml:"OBJID"`
	RESREAVESW string `xml:"RESREAVESW"`
	MULTICELLREBUILDSW string `xml:"MULTICELLREBUILDSW"`
}

