package model

import "encoding/xml"

type Nodebtrfoverloadthd struct {
	XMLName xml.Name `xml:"NODEBTRFOVERLOADTHD"`
	ATTRIBUTES NodebtrfoverloadthdAttributes `xml:"attributes"`
}

type NodebtrfoverloadthdAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	TRAFFICOVERLOADOTHD string `xml:"TRAFFICOVERLOADOTHD"`
	TRAFFICOVERLOADOPRD string `xml:"TRAFFICOVERLOADOPRD"`
	TRAFFICOVERLOADRTHD string `xml:"TRAFFICOVERLOADRTHD"`
	TRAFFICOVERLOADRPRD string `xml:"TRAFFICOVERLOADRPRD"`
	OBJID string `xml:"OBJID"`
}

