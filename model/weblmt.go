package model

import "encoding/xml"

type Weblmt struct {
	XMLName xml.Name `xml:"WEBLMT"`
	ATTRIBUTES WeblmtAttributes `xml:"attributes"`
}

type WeblmtAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	POLICY string `xml:"POLICY"`
	SSLVER string `xml:"SSLVER"`
}

