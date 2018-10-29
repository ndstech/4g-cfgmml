package model

import "encoding/xml"

type Vlanmap struct {
	XMLName xml.Name `xml:"VLANMAP"`
	ATTRIBUTES VlanmapAttributes `xml:"attributes"`
}

type VlanmapAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	VRFIDX string `xml:"VRFIDX"`
	NEXTHOPIP string `xml:"NEXTHOPIP"`
	MASK string `xml:"MASK"`
	VLANMODE string `xml:"VLANMODE"`
	VLANID string `xml:"VLANID"`
	VLANPRIO string `xml:"VLANPRIO"`
	SETPRIO string `xml:"SETPRIO"`
	VLANGROUPNO string `xml:"VLANGROUPNO"`
}

