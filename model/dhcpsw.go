package model

import "encoding/xml"

type Dhcpsw struct {
	XMLName xml.Name `xml:"DHCPSW"`
	ATTRIBUTES DhcpswAttributes `xml:"attributes"`
}

type DhcpswAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	SWITCH string `xml:"SWITCH"`
	VLANSCANSW string `xml:"VLANSCANSW"`
	DELAYSW string `xml:"DELAYSW"`
}

