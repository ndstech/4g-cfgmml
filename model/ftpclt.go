package model

import "encoding/xml"

type Ftpclt struct {
	XMLName xml.Name `xml:"FTPCLT"`
	ATTRIBUTES FtpcltAttributes `xml:"attributes"`
}

type FtpcltAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	ENCRYMODE string `xml:"ENCRYMODE"`
	SPTSTATEFWL string `xml:"SPTSTATEFWL"`
	SSLCERTAUTH string `xml:"SSLCERTAUTH"`
	MINDHLEN string `xml:"MINDHLEN"`
}

