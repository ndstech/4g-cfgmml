package model

import "encoding/xml"

type E1t1ber struct {
	XMLName xml.Name `xml:"E1T1BER"`
	ATTRIBUTES E1t1berAttributes `xml:"attributes"`
}

type E1t1berAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	BTL string `xml:"BTL"`
}

