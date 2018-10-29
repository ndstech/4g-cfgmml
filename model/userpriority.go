package model

import "encoding/xml"

type Userpriority struct {
	XMLName xml.Name `xml:"USERPRIORITY"`
	ATTRIBUTES UserpriorityAttributes `xml:"attributes"`
}

type UserpriorityAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	ArpSchSwitch string `xml:"ArpSchSwitch"`
	Arp1Priority string `xml:"Arp1Priority"`
	Arp2Priority string `xml:"Arp2Priority"`
	Arp3Priority string `xml:"Arp3Priority"`
	Arp4Priority string `xml:"Arp4Priority"`
	Arp5Priority string `xml:"Arp5Priority"`
	Arp6Priority string `xml:"Arp6Priority"`
	Arp7Priority string `xml:"Arp7Priority"`
	Arp8Priority string `xml:"Arp8Priority"`
	Arp9Priority string `xml:"Arp9Priority"`
	Arp10Priority string `xml:"Arp10Priority"`
	Arp11Priority string `xml:"Arp11Priority"`
	Arp12Priority string `xml:"Arp12Priority"`
	Arp13Priority string `xml:"Arp13Priority"`
	Arp14Priority string `xml:"Arp14Priority"`
	Arp15Priority string `xml:"Arp15Priority"`
	ObjId string `xml:"objId"`
}

