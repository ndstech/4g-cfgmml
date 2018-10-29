package model

import "encoding/xml"

type Dhcpsvrip struct {
	XMLName xml.Name `xml:"DHCPSVRIP"`
	ATTRIBUTES DhcpsvripAttributes `xml:"attributes"`
}

type DhcpsvripAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	DHCPSVRIP string `xml:"DHCPSVRIP"`
	DHCPRELAYIPSW string `xml:"DHCPRELAYIPSW"`
}

