package model

import "encoding/xml"

type Ippath struct {
	XMLName xml.Name `xml:"IPPATH"`
	ATTRIBUTES IppathAttributes `xml:"attributes"`
}

type IppathAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	PATHID string `xml:"PATHID"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	PT string `xml:"PT"`
	PN string `xml:"PN"`
	PATHTYPE string `xml:"PATHTYPE"`
	DSCP string `xml:"DSCP"`
	LOCALIP string `xml:"LOCALIP"`
	PEERIP string `xml:"PEERIP"`
	QT string `xml:"QT"`
	IPMUXSWITCH string `xml:"IPMUXSWITCH"`
	DESCRI string `xml:"DESCRI"`
	SBT string `xml:"SBT"`
	VRFIDX string `xml:"VRFIDX"`
	JNRSCGRP string `xml:"JNRSCGRP"`
	STATICCHK string `xml:"STATICCHK"`
}

