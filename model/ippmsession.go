package model

import "encoding/xml"

type Ippmsession struct {
	XMLName xml.Name `xml:"IPPMSESSION"`
	ATTRIBUTES IppmsessionAttributes `xml:"attributes"`
}

type IppmsessionAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	IPPMSN string `xml:"IPPMSN"`
	LOCALIP string `xml:"LOCALIP"`
	PEERIP string `xml:"PEERIP"`
	IPPMDSCP string `xml:"IPPMDSCP"`
	DIR string `xml:"DIR"`
	IPPMTYPE string `xml:"IPPMTYPE"`
	BINDPATH string `xml:"BINDPATH"`
	ASSOCIATEUPPATHPLR string `xml:"ASSOCIATEUPPATHPLR"`
	PATHID string `xml:"PATHID"`
}

