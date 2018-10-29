package model

import "encoding/xml"

type Synceth struct {
	XMLName xml.Name `xml:"SYNCETH"`
	ATTRIBUTES SyncethAttributes `xml:"attributes"`
}

type SyncethAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LN string `xml:"LN"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	PN string `xml:"PN"`
	QL string `xml:"QL"`
	SSM string `xml:"SSM"`
	PRI string `xml:"PRI"`
	TYPE string `xml:"TYPE"`
}

