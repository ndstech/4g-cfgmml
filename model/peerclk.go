package model

import "encoding/xml"

type Peerclk struct {
	XMLName xml.Name `xml:"PEERCLK"`
	ATTRIBUTES PeerclkAttributes `xml:"attributes"`
}

type PeerclkAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	PN string `xml:"PN"`
	PS string `xml:"PS"`
}

