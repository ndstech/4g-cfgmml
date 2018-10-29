package model

import "encoding/xml"

type Dhcprelayswitch struct {
	XMLName xml.Name `xml:"DHCPRELAYSWITCH"`
	ATTRIBUTES DhcprelayswitchAttributes `xml:"attributes"`
}

type DhcprelayswitchAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	ES string `xml:"ES"`
}

