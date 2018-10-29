package model

import "encoding/xml"

type Ikecfg struct {
	XMLName xml.Name `xml:"IKECFG"`
	ATTRIBUTES IkecfgAttributes `xml:"attributes"`
}

type IkecfgAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	IKELNM string `xml:"IKELNM"`
	IKEKLI string `xml:"IKEKLI"`
	IKEKLT string `xml:"IKEKLT"`
	NATKLI string `xml:"NATKLI"`
	DSCP string `xml:"DSCP"`
	IPSECSWITCHBACK string `xml:"IPSECSWITCHBACK"`
	IPSECSBWAITTIME string `xml:"IPSECSBWAITTIME"`
	IPSECSBRANDTIME string `xml:"IPSECSBRANDTIME"`
	IPSECSOWAITTIME string `xml:"IPSECSOWAITTIME"`
	IPSECSORANDTIME string `xml:"IPSECSORANDTIME"`
}

