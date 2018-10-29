package model

import "encoding/xml"

type Ipguard struct {
	XMLName xml.Name `xml:"IPGUARD"`
	ATTRIBUTES IpguardAttributes `xml:"attributes"`
}

type IpguardAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	ARPSPOOFCHKSW string `xml:"ARPSPOOFCHKSW"`
	ARPLRNSTRICTSW string `xml:"ARPLRNSTRICTSW"`
	INVALIDPKTCHKSW string `xml:"INVALIDPKTCHKSW"`
	IPSECREPLAYCHKSW string `xml:"IPSECREPLAYCHKSW"`
	REDIRECTDOSCHKSW string `xml:"REDIRECTDOSCHKSW"`
}

