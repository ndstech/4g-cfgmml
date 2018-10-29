package model

import "encoding/xml"

type Dscpmap struct {
	XMLName xml.Name `xml:"DSCPMAP"`
	ATTRIBUTES DscpmapAttributes `xml:"attributes"`
}

type DscpmapAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	VRFIDX string `xml:"VRFIDX"`
	DSCP string `xml:"DSCP"`
	VLANPRIO string `xml:"VLANPRIO"`
}

