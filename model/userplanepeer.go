package model

import "encoding/xml"

type Userplanepeer struct {
	XMLName xml.Name `xml:"USERPLANEPEER"`
	ATTRIBUTES UserplanepeerAttributes `xml:"attributes"`
}

type UserplanepeerAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	UPPEERID string `xml:"UPPEERID"`
	VRFIDX string `xml:"VRFIDX"`
	IPVERSION string `xml:"IPVERSION"`
	PEERIPV4 string `xml:"PEERIPV4"`
	IPSECSWITCH string `xml:"IPSECSWITCH"`
	REMOTEID string `xml:"REMOTEID"`
	CTRLMODE string `xml:"CTRLMODE"`
	AUTOCFGFLAG string `xml:"AUTOCFGFLAG"`
	USERLABEL string `xml:"USERLABEL"`
	STATICCHK string `xml:"STATICCHK"`
}

