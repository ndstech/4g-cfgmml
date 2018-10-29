package model

import "encoding/xml"

type Gtransparagw struct {
	XMLName xml.Name `xml:"GTRANSPARAGW"`
	ATTRIBUTES GtransparagwAttributes `xml:"attributes"`
}

type GtransparagwAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	PRIRULE string `xml:"PRIRULE"`
	LPSCHSW string `xml:"LPSCHSW"`
	RATECFGTYPE string `xml:"RATECFGTYPE"`
	SBTIME string `xml:"SBTIME"`
	ARPAGINGTIME string `xml:"ARPAGINGTIME"`
	MODE string `xml:"MODE"`
	OAMTYPE string `xml:"OAMTYPE"`
	IPERRFRMSW string `xml:"IPERRFRMSW"`
	FORWARDMODE string `xml:"FORWARDMODE"`
	SAMEPORTFORWARDSW string `xml:"SAMEPORTFORWARDSW"`
	PMAUTOSW string `xml:"PMAUTOSW"`
	IPCLKPRI string `xml:"IPCLKPRI"`
	PORTULOBSW string `xml:"PORTULOBSW"`
	PORTDLOBSW string `xml:"PORTDLOBSW"`
	PORTULCACSW string `xml:"PORTULCACSW"`
	PORTDLCACSW string `xml:"PORTDLCACSW"`
}

