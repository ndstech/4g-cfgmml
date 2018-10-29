package model

import "encoding/xml"

type Sctplnk struct {
	XMLName xml.Name `xml:"SCTPLNK"`
	ATTRIBUTES SctplnkAttributes `xml:"attributes"`
}

type SctplnkAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	SCTPNO string `xml:"SCTPNO"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	MAXSTREAM string `xml:"MAXSTREAM"`
	CTRLMODE string `xml:"CTRLMODE"`
	LOCIP string `xml:"LOCIP"`
	SECLOCIP string `xml:"SECLOCIP"`
	LOCPORT string `xml:"LOCPORT"`
	PEERIP string `xml:"PEERIP"`
	SECPEERIP string `xml:"SECPEERIP"`
	PEERPORT string `xml:"PEERPORT"`
	RTOMIN string `xml:"RTOMIN"`
	RTOMAX string `xml:"RTOMAX"`
	RTOINIT string `xml:"RTOINIT"`
	RTOALPHA string `xml:"RTOALPHA"`
	RTOBETA string `xml:"RTOBETA"`
	HBINTER string `xml:"HBINTER"`
	MAXASSOCRETR string `xml:"MAXASSOCRETR"`
	MAXPATHRETR string `xml:"MAXPATHRETR"`
	CHKSUMTYPE string `xml:"CHKSUMTYPE"`
	AUTOSWITCH string `xml:"AUTOSWITCH"`
	SWITCHBACKHBNUM string `xml:"SWITCHBACKHBNUM"`
	BLKFLAG string `xml:"BLKFLAG"`
	TSACK string `xml:"TSACK"`
	DESCRI string `xml:"DESCRI"`
	AUTOCFGFLAG string `xml:"AUTOCFGFLAG"`
	VRFIDX string `xml:"VRFIDX"`
}

