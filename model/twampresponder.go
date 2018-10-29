package model

import "encoding/xml"

type Twampresponder struct {
	XMLName xml.Name `xml:"TWAMPRESPONDER"`
	ATTRIBUTES TwampresponderAttributes `xml:"attributes"`
}

type TwampresponderAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	RESPONDERID string `xml:"RESPONDERID"`
	LOCALIP string `xml:"LOCALIP"`
	VRFINDEX string `xml:"VRFINDEX"`
	DSCP string `xml:"DSCP"`
	REFWAIT string `xml:"REFWAIT"`
	SERVWAIT string `xml:"SERVWAIT"`
	LOCALPORT string `xml:"LOCALPORT"`
	TWAMPARCH string `xml:"TWAMPARCH"`
}

