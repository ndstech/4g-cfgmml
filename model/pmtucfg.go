package model

import "encoding/xml"

type Pmtucfg struct {
	XMLName xml.Name `xml:"PMTUCFG"`
	ATTRIBUTES PmtucfgAttributes `xml:"attributes"`
}

type PmtucfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	MODE string `xml:"MODE"`
	UDPPORT string `xml:"UDPPORT"`
	DETECTPERIOD string `xml:"DETECTPERIOD"`
	AGINGPERIOD string `xml:"AGINGPERIOD"`
	DSCP string `xml:"DSCP"`
}

