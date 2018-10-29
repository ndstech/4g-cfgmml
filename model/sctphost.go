package model

import "encoding/xml"

type Sctphost struct {
	XMLName xml.Name `xml:"SCTPHOST"`
	ATTRIBUTES SctphostAttributes `xml:"attributes"`
}

type SctphostAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	SCTPHOSTID string `xml:"SCTPHOSTID"`
	VRFIDX string `xml:"VRFIDX"`
	IPVERSION string `xml:"IPVERSION"`
	SIGIP1V4 string `xml:"SIGIP1V4"`
	SIGIP1SECSWITCH string `xml:"SIGIP1SECSWITCH"`
	SIGIP2V4 string `xml:"SIGIP2V4"`
	SIGIP2SECSWITCH string `xml:"SIGIP2SECSWITCH"`
	PN string `xml:"PN"`
	SCTPTEMPLATEID string `xml:"SCTPTEMPLATEID"`
	USERLABEL string `xml:"USERLABEL"`
	SIMPLEMODESWITCH string `xml:"SIMPLEMODESWITCH"`
}

