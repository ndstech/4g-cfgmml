package model

import "encoding/xml"

type Htcdpa struct {
	XMLName xml.Name `xml:"HTCDPA"`
	ATTRIBUTES HtcdpaAttributes `xml:"attributes"`
}

type HtcdpaAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CN string `xml:"CN"`
	SRN string `xml:"SRN"`
	SN string `xml:"SN"`
	LTCP string `xml:"LTCP"`
	HTCP string `xml:"HTCP"`
	TLT string `xml:"TLT"`
	DBD string `xml:"DBD"`
	NTDI string `xml:"NTDI"`
	NTDO string `xml:"NTDO"`
	HTDO string `xml:"HTDO"`
}

