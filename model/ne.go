package model

import "encoding/xml"

type Ne struct {
	XMLName    xml.Name     `xml:"NE"`
	ATTRIBUTES NeAttributes `xml:"attributes"`
}

type NeAttributes struct {
	XMLName         xml.Name `xml:"attributes"`
	LOCATION        string   `xml:"LOCATION"`
	SWVERSION       string   `xml:"SWVERSION"`
	USERLABEL       string   `xml:"USERLABEL"`
	NERMVERSION     string   `xml:"NERMVERSION"`
	INTERFACEID     string   `xml:"INTERFACEID"`
	PRODUCTVERSION  string   `xml:"PRODUCTVERSION"`
	LMTVERSION      string   `xml:"LMTVERSION"`
	DID             string   `xml:"DID"`
	SITENAME        string   `xml:"SITENAME"`
	NENAME          string   `xml:"NENAME"`
	HOTPATCHVERSION string   `xml:"HOTPATCHVERSION"`
	CLOUDBBID       string   `xml:"CLOUDBBID"`
}
