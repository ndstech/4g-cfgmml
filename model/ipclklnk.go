package model

import "encoding/xml"

type Ipclklnk struct {
	XMLName xml.Name `xml:"IPCLKLNK"`
	ATTRIBUTES IpclklnkAttributes `xml:"attributes"`
}

type IpclklnkAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	LN string `xml:"LN"`
	ICPT string `xml:"ICPT"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	CIP string `xml:"CIP"`
	SIP string `xml:"SIP"`
	DM string `xml:"DM"`
	DELAYTYPE string `xml:"DELAYTYPE"`
	CLASSIDENTIFY string `xml:"CLASSIDENTIFY"`
	CLASS0 string `xml:"CLASS0"`
	CLASS1 string `xml:"CLASS1"`
	CLASS2 string `xml:"CLASS2"`
	CLASS3 string `xml:"CLASS3"`
	CNM string `xml:"CNM"`
	CMPST string `xml:"CMPST"`
	PRI string `xml:"PRI"`
	ANNFREQ string `xml:"ANNFREQ"`
	SYNCFREQ string `xml:"SYNCFREQ"`
	MASTERPRIO string `xml:"MASTERPRIO"`
	PROFILETYPE string `xml:"PROFILETYPE"`
	NEGDURATION string `xml:"NEGDURATION"`
	IPMODE string `xml:"IPMODE"`
	PREHOPIP string `xml:"PREHOPIP"`
	PDELAYREQFREQ string `xml:"PDELAYREQFREQ"`
	DEVTYPE string `xml:"DEVTYPE"`
	DSTMLTMACTYPE string `xml:"DSTMLTMACTYPE"`
	DELAYFREQ string `xml:"DELAYFREQ"`
	IPSYNCMODE string `xml:"IPSYNCMODE"`
}

