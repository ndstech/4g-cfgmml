package model

import "encoding/xml"

type Crlpolicy struct {
	XMLName xml.Name `xml:"CRLPOLICY"`
	ATTRIBUTES CrlpolicyAttributes `xml:"attributes"`
}

type CrlpolicyAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	CRLPOLICY string `xml:"CRLPOLICY"`
}

