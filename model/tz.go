package model

import "encoding/xml"

type Tz struct {
	XMLName xml.Name `xml:"TZ"`
	ATTRIBUTES TzAttributes `xml:"attributes"`
}

type TzAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	ZONET string `xml:"ZONET"`
	DST string `xml:"DST"`
}

