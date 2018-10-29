package model

import "encoding/xml"

type Ulocellrsclmtpara struct {
	XMLName xml.Name `xml:"ULOCELLRSCLMTPARA"`
	ATTRIBUTES UlocellrsclmtparaAttributes `xml:"attributes"`
}

type UlocellrsclmtparaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	ULOCELLID string `xml:"ULOCELLID"`
	CME8KRSCLMT string `xml:"CME8KRSCLMT"`
	CME16KRSCLMT string `xml:"CME16KRSCLMT"`
	CME32KRSCLMT string `xml:"CME32KRSCLMT"`
	CME64KRSCLMT string `xml:"CME64KRSCLMT"`
	CME128KRSCLMT string `xml:"CME128KRSCLMT"`
	CME256KRSCLMT string `xml:"CME256KRSCLMT"`
	CME384KRSCLMT string `xml:"CME384KRSCLMT"`
	CME512KRSCLMT string `xml:"CME512KRSCLMT"`
	CME768KRSCLMT string `xml:"CME768KRSCLMT"`
	CME1024KRSCLMT string `xml:"CME1024KRSCLMT"`
	CME1536KRSCLMT string `xml:"CME1536KRSCLMT"`
	CME1800KRSCLMT string `xml:"CME1800KRSCLMT"`
}

