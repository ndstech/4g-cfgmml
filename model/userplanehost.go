package model

import "encoding/xml"

type Userplanehost struct {
	XMLName xml.Name `xml:"USERPLANEHOST"`
	ATTRIBUTES UserplanehostAttributes `xml:"attributes"`
}

type UserplanehostAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	UPHOSTID string `xml:"UPHOSTID"`
	VRFIDX string `xml:"VRFIDX"`
	IPVERSION string `xml:"IPVERSION"`
	LOCIPV4 string `xml:"LOCIPV4"`
	IPSECSWITCH string `xml:"IPSECSWITCH"`
	USERLABEL string `xml:"USERLABEL"`
}

