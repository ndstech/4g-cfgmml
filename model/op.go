package model

import "encoding/xml"

type Op struct {
	XMLName xml.Name `xml:"OP"`
	ATTRIBUTES OpAttributes `xml:"attributes"`
}

type OpAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	SWOP string `xml:"SWOP"`
	LOCKST string `xml:"LOCKST"`
}

