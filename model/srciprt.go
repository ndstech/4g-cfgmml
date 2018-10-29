package model

import "encoding/xml"

type Srciprt struct {
	XMLName xml.Name `xml:"SRCIPRT"`
	ATTRIBUTES SrciprtAttributes `xml:"attributes"`
}

type SrciprtAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	SRCRTIDX string `xml:"SRCRTIDX"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	SBT string `xml:"SBT"`
	SRCIP string `xml:"SRCIP"`
	RTTYPE string `xml:"RTTYPE"`
	NEXTHOP string `xml:"NEXTHOP"`
	PREF string `xml:"PREF"`
	USERLABEL string `xml:"USERLABEL"`
}

