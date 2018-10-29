package model

import "encoding/xml"

type Sctppeer struct {
	XMLName xml.Name `xml:"SCTPPEER"`
	ATTRIBUTES SctppeerAttributes `xml:"attributes"`
}

type SctppeerAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	SCTPPEERID string `xml:"SCTPPEERID"`
	VRFIDX string `xml:"VRFIDX"`
	IPVERSION string `xml:"IPVERSION"`
	SIGIP1V4 string `xml:"SIGIP1V4"`
	SIGIP1SECSWITCH string `xml:"SIGIP1SECSWITCH"`
	SIGIP2V4 string `xml:"SIGIP2V4"`
	SIGIP2SECSWITCH string `xml:"SIGIP2SECSWITCH"`
	PN string `xml:"PN"`
	REMOTEID string `xml:"REMOTEID"`
	CTRLMODE string `xml:"CTRLMODE"`
	AUTOCFGFLAG string `xml:"AUTOCFGFLAG"`
	USERLABEL string `xml:"USERLABEL"`
	BLKFLAG string `xml:"BLKFLAG"`
	SIMPLEMODESWITCH string `xml:"SIMPLEMODESWITCH"`
}

