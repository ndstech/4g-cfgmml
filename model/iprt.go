package model

import "encoding/xml"

type Iprt struct {
	XMLName xml.Name `xml:"IPRT"`
	ATTRIBUTES IprtAttributes `xml:"attributes"`
}

type IprtAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	RTIDX string `xml:"RTIDX"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	SBT string `xml:"SBT"`
	RTTYPE string `xml:"RTTYPE"`
	VRFIDX string `xml:"VRFIDX"`
	DSTIP string `xml:"DSTIP"`
	DSTMASK string `xml:"DSTMASK"`
	NEXTHOP string `xml:"NEXTHOP"`
	MTUSWITCH string `xml:"MTUSWITCH"`
	PREF string `xml:"PREF"`
	DESCRI string `xml:"DESCRI"`
}

