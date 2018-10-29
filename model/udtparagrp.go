package model

import "encoding/xml"

type Udtparagrp struct {
	XMLName xml.Name `xml:"UDTPARAGRP"`
	ATTRIBUTES UdtparagrpAttributes `xml:"attributes"`
}

type UdtparagrpAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	UDTPARAGRPID string `xml:"UDTPARAGRPID"`
	PRIRULE string `xml:"PRIRULE"`
	PRI string `xml:"PRI"`
	ACTFACTOR string `xml:"ACTFACTOR"`
	PRIMTRANRSCTYPE string `xml:"PRIMTRANRSCTYPE"`
	PRIMPTLOADTH string `xml:"PRIMPTLOADTH"`
	PRIM2SECPTLOADRATH string `xml:"PRIM2SECPTLOADRATH"`
}

