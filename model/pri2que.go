package model

import "encoding/xml"

type Pri2que struct {
	XMLName xml.Name `xml:"PRI2QUE"`
	ATTRIBUTES Pri2queAttributes `xml:"attributes"`
}

type Pri2queAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	VRFIDX string `xml:"VRFIDX"`
	PRI0 string `xml:"PRI0"`
	PRI1 string `xml:"PRI1"`
	PRI2 string `xml:"PRI2"`
	PRI3 string `xml:"PRI3"`
	PRI4 string `xml:"PRI4"`
	PRI5 string `xml:"PRI5"`
	PRI6 string `xml:"PRI6"`
}

