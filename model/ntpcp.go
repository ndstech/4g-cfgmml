package model

import "encoding/xml"

type Ntpcp struct {
	XMLName xml.Name `xml:"NTPCP"`
	ATTRIBUTES NtpcpAttributes `xml:"attributes"`
}

type NtpcpAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	IP string `xml:"IP"`
	SYNCCYCLE string `xml:"SYNCCYCLE"`
	PORT string `xml:"PORT"`
	MODE string `xml:"MODE"`
	AUTHMODE string `xml:"AUTHMODE"`
	KEY string `xml:"KEY"`
	KEYID string `xml:"KEYID"`
	MASTERFLAG string `xml:"MASTERFLAG"`
}

