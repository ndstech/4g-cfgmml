package model

import "encoding/xml"

type Node struct {
	XMLName xml.Name `xml:"NODE"`
	ATTRIBUTES NodeAttributes `xml:"attributes"`
}

type NodeAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	PRODUCTTYPE string `xml:"PRODUCTTYPE"`
	USERLABEL string `xml:"USERLABEL"`
	NERMVERSION string `xml:"NERMVERSION"`
	NODEID string `xml:"NODEID"`
	NODENAME string `xml:"NODENAME"`
	WM string `xml:"WM"`
	SWVERSION string `xml:"SWVERSION"`
	HOTPATCHVERSION string `xml:"HOTPATCHVERSION"`
	PRODUCTVERSION string `xml:"PRODUCTVERSION"`
	INTERFACEID string `xml:"INTERFACEID"`
	LMTVERSION string `xml:"LMTVERSION"`
	APPVERSION string `xml:"APPVERSION"`
	APPHOTPATCHVERSION string `xml:"APPHOTPATCHVERSION"`
	LTEMODE string `xml:"LTEMODE"`
}

