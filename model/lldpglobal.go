package model

import "encoding/xml"

type Lldpglobal struct {
	XMLName xml.Name `xml:"LLDPGLOBAL"`
	ATTRIBUTES LldpglobalAttributes `xml:"attributes"`
}

type LldpglobalAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	TXINTVAL string `xml:"TXINTVAL"`
	HOLDMULTI string `xml:"HOLDMULTI"`
	REINITDELAY string `xml:"REINITDELAY"`
	DELAY string `xml:"DELAY"`
	NOTIFYSW string `xml:"NOTIFYSW"`
	NOTIFYINTERVAL string `xml:"NOTIFYINTERVAL"`
}

